package entry

import (
	"context"
	"errors"
	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
	"time"

	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/google/uuid"

	"github.com/asaskevich/govalidator"
)

type EntryUseCase interface {
	http
	async
}

type http interface {
	CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error)
	DescribeEntry(ctx context.Context, input DescribeEntryInput) (*DescribeEntryOutput, error)
	ListEntries(ctx context.Context, input ListEntryInput) (*ListEntryOutput, error)
	UpdateEntry(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error)
	RemoveEntry(ctx context.Context, input RemoveEntryInput) (*RemoveEntryOutput, error)
	PublishEntry(ctx context.Context, input PublishEntryInput) (*PublishEntryOutput, error)
}

type async interface {
	EntryCreatedEvent(ctx context.Context, input EntryCreatedEventInput) error
}

func NewDefaultEntryUseCase(options DefaultEntryUseCaseOptions) (*DefaultEntryUseCase, error) {
	if options.EntryStorage == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided EntryStorage is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided ContaBlockUseCase is nil")
	}
	if options.BalanceUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided BalanceUseCase is nil")
	}
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided AccountUseCase is nil")
	}
	if options.ContractUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided ContractUseCase is nil")
	}
	if options.UserUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided UserUseCase is nil")
	}

	return &DefaultEntryUseCase{
		entryStorage:      options.EntryStorage,
		contaBlockUseCase: options.ContaBlockUseCase,
		balanceUseCase:    options.BalanceUseCase,
		accountUseCase:    options.AccountUseCase,
		contractUseCase:   options.ContractUseCase,
		userUseCase:       options.UserUseCase,
	}, nil
}

type DefaultEntryUseCase struct {
	entryStorage      EntryStorage
	contaBlockUseCase contablock.ContaBlockUseCase
	balanceUseCase    balance.BalanceUseCase
	accountUseCase    account.AccountUseCase
	contractUseCase   contract.ContractUseCase
	userUseCase       user.UserUseCase
}

type DefaultEntryUseCaseOptions struct {
	EntryStorage      EntryStorage
	ContaBlockUseCase contablock.ContaBlockUseCase
	BalanceUseCase    balance.BalanceUseCase
	AccountUseCase    account.AccountUseCase
	ContractUseCase   contract.ContractUseCase
	UserUseCase       user.UserUseCase
}

var _ EntryUseCase = (*DefaultEntryUseCase)(nil)

func (useCase *DefaultEntryUseCase) CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	if input.DebitBalance == nil && input.CreditBalance == nil {
		return nil, errors.New("valid debit or credit balance is required")
	}

	if input.DebitBalance != nil && input.CreditBalance != nil {
		return nil, errors.New("only one of debit or credit balance is required")
	}

	id := uuid.New().String()

	var describeAccountInput account.DescribeAccountInput
	describeAccountInput.Project = input.Project
	describeAccountInput.ID = input.AccountID
	storedAccount, err := useCase.accountUseCase.DescribeAccount(ctx, describeAccountInput)
	if err != nil {
		return nil, err
	}

	var balanceByYearID *string
	if storedAccount.Account.BalanceID != nil && *storedAccount.Account.BalanceID != "" {
		var createBalanceInput balance.CreateBalanceByYearInput
		createBalanceInput.Project = input.Project
		createBalanceInput.BalanceID = *storedAccount.Account.BalanceID
		createBalanceInput.Amount = *getAmountForBalance(input.DebitBalance, input.CreditBalance)
		createBalanceInput.Year = input.Date.Year()
		storedBalance, err := useCase.balanceUseCase.CreateBalanceByYear(ctx, createBalanceInput)
		if err != nil {
			return nil, err
		}
		balanceByYearID = &storedBalance.ID
	}

	var desiredEntry EntryStorageData
	if input.ID == nil {
		input.ID = &id
	}
	desiredEntry.Project = input.Project
	desiredEntry.ID = *input.ID
	desiredEntry.ContractID = input.ContractID
	desiredEntry.UserID = input.UserID
	desiredEntry.AccountID = input.AccountID
	desiredEntry.MultipleEntryID = input.MultipleEntryID
	desiredEntry.BalanceByYearID = balanceByYearID
	desiredEntry.Currency = input.Currency
	desiredEntry.Concept = input.Concept
	desiredEntry.Date = input.Date
	quarter := useCase.calculateQuarter(input.Date)
	desiredEntry.Quarter = quarter
	if input.DebitBalance == nil {
		debitBalance := 0.00
		input.DebitBalance = &debitBalance
	}
	desiredEntry.DebitBalance = *input.DebitBalance
	if input.CreditBalance == nil {
		creditBalance := 0.00
		input.CreditBalance = &creditBalance
	}
	desiredEntry.CreditBalance = *input.CreditBalance
	desiredEntry.TotalBalance = desiredEntry.DebitBalance - desiredEntry.CreditBalance
	desiredEntry.State = EntryStateCreated
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	storedEntry, err := useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	// Call SC
	var describeContractInput contract.DescribeContractInput
	describeContractInput.Project = input.Project
	describeContractInput.ID = input.ContractID
	storedContract, err := useCase.contractUseCase.DescribeContract(ctx, describeContractInput)
	if err != nil {
		return nil, err
	}

	var describeUserInput user.DescribeUserInput
	describeUserInput.Project = input.Project
	describeUserInput.ID = input.UserID
	storedUser, err := useCase.userUseCase.DescribeUser(ctx, describeUserInput)
	if err != nil {
		return nil, err
	}

	var createEntryInput contablock.CreateEntryInput
	createEntryInput.ContractAddress = storedContract.Contract.EthAddress
	createEntryInput.UserAddress = storedUser.User.EthAddress
	createEntryInput.UserPrivateKey = storedUser.User.PrivateKey
	createEntryInput.ID = id
	createEntryInput.ContractID = input.ContractID
	createEntryInput.UserID = input.UserID
	createEntryInput.AccountID = input.AccountID
	createEntryInput.MultipleEntryID = input.MultipleEntryID
	createEntryInput.Currency = input.Currency
	if input.Concept != nil {
		createEntryInput.Concept = *input.Concept
	}
	createEntryInput.Date = input.Date
	createEntryInput.DebitBalance = input.DebitBalance
	createEntryInput.CreditBalance = input.CreditBalance
	err = useCase.contaBlockUseCase.CreateEntry(createEntryInput)
	if err != nil {
		return nil, err
	}

	var output CreateEntryOutput
	output.Entry.Project = storedEntry.Project
	output.Entry.ID = storedEntry.ID
	output.Entry.ContractID = storedEntry.ContractID
	output.Entry.UserID = storedEntry.UserID
	output.Entry.AccountID = storedEntry.AccountID
	output.Entry.MultipleEntryID = storedEntry.MultipleEntryID
	output.Entry.BalanceByYearID = storedEntry.BalanceByYearID
	output.Entry.Currency = storedEntry.Currency
	output.Entry.Concept = storedEntry.Concept
	output.Entry.Date = storedEntry.Date
	output.Entry.Quarter = storedEntry.Quarter
	output.Entry.DebitBalance = storedEntry.DebitBalance
	output.Entry.CreditBalance = storedEntry.CreditBalance
	output.Entry.TotalBalance = storedEntry.TotalBalance
	output.Entry.State = storedEntry.State
	output.Entry.StateReason = storedEntry.StateReason
	output.Entry.CreationDate = storedEntry.CreationDate
	output.Entry.LastUpdate = storedEntry.LastUpdate

	return &output, nil
}

func (useCase *DefaultEntryUseCase) DescribeEntry(ctx context.Context, input DescribeEntryInput) (*DescribeEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedEntry, err := useCase.entryStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeEntryOutput
	output.Entry.Project = storedEntry.Project
	output.Entry.ID = storedEntry.ID
	output.Entry.ContractID = storedEntry.ContractID
	output.Entry.UserID = storedEntry.UserID
	output.Entry.AccountID = storedEntry.AccountID
	output.Entry.MultipleEntryID = storedEntry.MultipleEntryID
	output.Entry.BalanceByYearID = storedEntry.BalanceByYearID
	output.Entry.Currency = storedEntry.Currency
	output.Entry.Concept = storedEntry.Concept
	output.Entry.Date = storedEntry.Date
	output.Entry.Quarter = storedEntry.Quarter
	output.Entry.DebitBalance = storedEntry.DebitBalance
	output.Entry.CreditBalance = storedEntry.CreditBalance
	output.Entry.TotalBalance = storedEntry.TotalBalance
	output.Entry.State = storedEntry.State
	output.Entry.StateReason = storedEntry.StateReason
	output.Entry.CreationDate = storedEntry.CreationDate
	output.Entry.LastUpdate = storedEntry.LastUpdate

	return &output, err
}

func (useCase *DefaultEntryUseCase) ListEntries(ctx context.Context, input ListEntryInput) (*ListEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var filters Filters
	filters.MultipleEntryID = input.MultipleEntryID
	filters.AccountID = input.AccountID
	storedEntries, err := useCase.entryStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	var items []Entry
	for _, item := range storedEntries.Items {
		var entry Entry
		entry.Project = item.Project
		entry.ID = item.ID
		entry.ContractID = item.ContractID
		entry.UserID = item.UserID
		entry.AccountID = item.AccountID
		entry.MultipleEntryID = item.MultipleEntryID
		entry.BalanceByYearID = item.BalanceByYearID
		entry.Currency = item.Currency
		entry.Concept = item.Concept
		entry.Date = item.Date
		entry.Quarter = item.Quarter
		entry.DebitBalance = item.DebitBalance
		entry.CreditBalance = item.CreditBalance
		entry.TotalBalance = item.TotalBalance
		entry.State = item.State
		entry.StateReason = item.StateReason
		entry.CreationDate = item.CreationDate
		entry.LastUpdate = item.LastUpdate

		items = append(items, entry)
	}

	return &ListEntryOutput{
		Items: items,
	}, err
}

func (useCase *DefaultEntryUseCase) UpdateEntry(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.Project = input.Entry.Project
	describeEntryInput.ID = input.Entry.ID

	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	var desiredEntry EntryStorageData
	desiredEntry.Project = input.Entry.Project
	desiredEntry.ID = input.Entry.ID
	desiredEntry.ContractID = input.Entry.ContractID
	desiredEntry.UserID = input.Entry.UserID
	desiredEntry.AccountID = input.Entry.AccountID
	desiredEntry.MultipleEntryID = input.Entry.MultipleEntryID
	desiredEntry.BalanceByYearID = input.Entry.BalanceByYearID
	desiredEntry.Currency = input.Entry.Currency
	if input.Entry.Concept == nil {
		desiredEntry.Concept = storedEntry.Entry.Concept
	} else {
		desiredEntry.Concept = input.Entry.Concept
	}
	desiredEntry.Date = input.Entry.Date
	quarter := useCase.calculateQuarter(input.Entry.Date)
	desiredEntry.Quarter = quarter
	desiredEntry.DebitBalance = input.Entry.DebitBalance
	desiredEntry.CreditBalance = input.Entry.CreditBalance
	desiredEntry.TotalBalance = desiredEntry.DebitBalance - desiredEntry.CreditBalance
	desiredEntry.State = input.Entry.State
	desiredEntry.CreationDate = storedEntry.Entry.CreationDate
	desiredEntry.LastUpdate = time.Now()

	updatedEntry, err := useCase.entryStorage.Edit(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var output UpdateEntryOutput
	output.Entry.Project = updatedEntry.Project
	output.Entry.ID = updatedEntry.ID
	output.Entry.ContractID = updatedEntry.ContractID
	output.Entry.UserID = updatedEntry.UserID
	output.Entry.AccountID = updatedEntry.AccountID
	output.Entry.MultipleEntryID = updatedEntry.MultipleEntryID
	output.Entry.BalanceByYearID = updatedEntry.BalanceByYearID
	output.Entry.Currency = updatedEntry.Currency
	output.Entry.Concept = updatedEntry.Concept
	output.Entry.Date = updatedEntry.Date
	output.Entry.Quarter = updatedEntry.Quarter
	output.Entry.DebitBalance = updatedEntry.DebitBalance
	output.Entry.CreditBalance = updatedEntry.CreditBalance
	output.Entry.TotalBalance = updatedEntry.TotalBalance
	output.Entry.State = updatedEntry.State
	output.Entry.StateReason = updatedEntry.StateReason
	output.Entry.CreationDate = updatedEntry.CreationDate
	output.Entry.LastUpdate = updatedEntry.LastUpdate

	return &output, err
}

func (useCase *DefaultEntryUseCase) RemoveEntry(ctx context.Context, input RemoveEntryInput) (*RemoveEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.ID
	describeEntryInput.Project = input.Project
	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	// Call SC
	var describeContractInput contract.DescribeContractInput
	describeContractInput.Project = input.Project
	describeContractInput.ID = input.ContractID
	storedContract, err := useCase.contractUseCase.DescribeContract(ctx, describeContractInput)
	if err != nil {
		return nil, err
	}

	var describeUserInput user.DescribeUserInput
	describeUserInput.Project = input.Project
	describeUserInput.ID = input.UserID
	storedUser, err := useCase.userUseCase.DescribeUser(ctx, describeUserInput)
	if err != nil {
		return nil, err
	}

	var getEntryInput contablock.GetEntryDataInput
	getEntryInput.ID = input.ID
	getEntryInput.ContractAddress = storedContract.Contract.EthAddress
	getEntryInput.UserAddress = storedUser.User.EthAddress
	getEntryInput.UserPrivateKey = storedUser.User.PrivateKey
	entryData, err := useCase.contaBlockUseCase.GetEntryData(getEntryInput)
	if err != nil {
		return nil, err
	}

	if !(storedEntry.Entry.State == EntryStateCreated && entryData.State == contablock.EntryStateCreated) {
		return nil, errors.New("failure publishing state not allowed")
	}

	if storedEntry.Entry.BalanceByYearID != nil && *storedEntry.Entry.BalanceByYearID != "" {
		var removeBalanceInput balance.RemoveBalanceInput
		removeBalanceInput.Project = input.Project
		removeBalanceInput.ID = *storedEntry.Entry.BalanceByYearID
		_, err = useCase.balanceUseCase.RemoveBalance(ctx, removeBalanceInput)
		if err != nil {
			return nil, err
		}
	}

	var removeEntrySmartContractInput contablock.RemoveEntryInput
	removeEntrySmartContractInput.ID = input.ID
	removeEntrySmartContractInput.ContractAddress = storedContract.Contract.EthAddress
	removeEntrySmartContractInput.UserAddress = storedUser.User.EthAddress
	removeEntrySmartContractInput.UserPrivateKey = storedUser.User.PrivateKey
	err = useCase.contaBlockUseCase.RemoveEntry(removeEntrySmartContractInput)
	if err != nil {
		return nil, err
	}

	removedEntry, err := useCase.entryStorage.Remove(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output RemoveEntryOutput
	output.Entry.Project = removedEntry.Project
	output.Entry.ID = removedEntry.ID
	output.Entry.ContractID = removedEntry.ContractID
	output.Entry.UserID = removedEntry.UserID
	output.Entry.AccountID = removedEntry.AccountID
	output.Entry.MultipleEntryID = removedEntry.MultipleEntryID
	output.Entry.BalanceByYearID = removedEntry.BalanceByYearID
	output.Entry.Currency = removedEntry.Currency
	output.Entry.Concept = removedEntry.Concept
	output.Entry.Date = removedEntry.Date
	output.Entry.Quarter = removedEntry.Quarter
	output.Entry.DebitBalance = removedEntry.DebitBalance
	output.Entry.CreditBalance = removedEntry.CreditBalance
	output.Entry.TotalBalance = removedEntry.TotalBalance
	output.Entry.State = removedEntry.State
	output.Entry.StateReason = removedEntry.StateReason
	output.Entry.CreationDate = removedEntry.CreationDate
	output.Entry.LastUpdate = removedEntry.LastUpdate

	return &output, err
}

func (useCase *DefaultEntryUseCase) PublishEntry(ctx context.Context, input PublishEntryInput) (*PublishEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.Project = input.Project
	describeEntryInput.ID = input.ID
	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	// Call SC
	var describeContractInput contract.DescribeContractInput
	describeContractInput.Project = input.Project
	describeContractInput.ID = input.ContractID
	storedContract, err := useCase.contractUseCase.DescribeContract(ctx, describeContractInput)
	if err != nil {
		return nil, err
	}

	var describeUserInput user.DescribeUserInput
	describeUserInput.Project = input.Project
	describeUserInput.ID = input.UserID
	storedUser, err := useCase.userUseCase.DescribeUser(ctx, describeUserInput)
	if err != nil {
		return nil, err
	}

	var getEntryInput contablock.GetEntryDataInput
	getEntryInput.ID = input.ID
	getEntryInput.ContractAddress = storedContract.Contract.EthAddress
	getEntryInput.UserAddress = storedUser.User.EthAddress
	getEntryInput.UserPrivateKey = storedUser.User.PrivateKey
	entryData, err := useCase.contaBlockUseCase.GetEntryData(getEntryInput)
	if err != nil {
		return nil, err
	}

	if !(storedEntry.Entry.State == EntryStateCreated && entryData.State == contablock.EntryStateCreated) {
		return nil, errors.New("failure publishing state not allowed")
	}

	var publishEntryInput contablock.PublishEntryInput
	publishEntryInput.ID = input.ID
	publishEntryInput.ContractAddress = storedContract.Contract.EthAddress
	publishEntryInput.UserAddress = storedUser.User.EthAddress
	publishEntryInput.UserPrivateKey = storedUser.User.PrivateKey
	err = useCase.contaBlockUseCase.PublishEntry(publishEntryInput)
	if err != nil {
		return nil, err
	}

	var updateEntryInput UpdateEntryInput
	updateEntryInput.Entry = storedEntry.Entry
	updateEntryInput.Entry.State = EntryStatePublished
	updatedEntry, err := useCase.UpdateEntry(ctx, updateEntryInput)
	if err != nil {
		return nil, err
	}

	var output PublishEntryOutput
	output.Entry = updatedEntry.Entry

	return &output, err
}

func (useCase *DefaultEntryUseCase) EntryCreatedEvent(ctx context.Context, input EntryCreatedEventInput) error {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return errors.New("invalid input: " + err.Error())
	}

	var describeAccountInput account.DescribeAccountInput
	describeAccountInput.Project = input.Project
	describeAccountInput.ID = input.AccountID
	storedAccount, err := useCase.accountUseCase.DescribeAccount(ctx, describeAccountInput)
	if err != nil {
		return err
	}

	var balanceByYearID *string
	if storedAccount.Account.BalanceID != nil && *storedAccount.Account.BalanceID != "" {
		var createBalanceInput balance.CreateBalanceByYearInput
		createBalanceInput.Project = input.Project
		createBalanceInput.BalanceID = *storedAccount.Account.BalanceID
		createBalanceInput.Amount = *getAmountForBalance(input.DebitBalance, input.CreditBalance)
		createBalanceInput.Year = input.Date.Year()
		storedBalance, err := useCase.balanceUseCase.CreateBalanceByYear(ctx, createBalanceInput)
		if err != nil {
			return err
		}
		balanceByYearID = &storedBalance.ID
	}

	var desiredEntry EntryStorageData
	desiredEntry.Project = input.Project
	desiredEntry.ID = input.ID
	desiredEntry.ContractID = input.ContractID
	desiredEntry.UserID = input.UserID
	desiredEntry.AccountID = input.AccountID
	desiredEntry.MultipleEntryID = input.MultipleEntryID
	desiredEntry.BalanceByYearID = balanceByYearID
	desiredEntry.Currency = input.Currency
	desiredEntry.Concept = &input.Concept
	desiredEntry.Date = input.Date
	quarter := useCase.calculateQuarter(input.Date)
	desiredEntry.Quarter = quarter
	if input.DebitBalance == nil {
		debitBalance := 0.00
		input.DebitBalance = &debitBalance
	}
	desiredEntry.DebitBalance = *input.DebitBalance
	if input.CreditBalance == nil {
		creditBalance := 0.00
		input.CreditBalance = &creditBalance
	}
	desiredEntry.CreditBalance = *input.CreditBalance
	desiredEntry.TotalBalance = desiredEntry.DebitBalance - desiredEntry.CreditBalance
	desiredEntry.State = EntryStateCreated
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	_, err = useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *DefaultEntryUseCase) calculateQuarter(date time.Time) int {
	month := date.Month()
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		return 0
	}
}

func getAmountForBalance(debitBalance *float64, creditBalance *float64) *float64 {
	if debitBalance == nil {
		return creditBalance
	}

	return debitBalance
}
