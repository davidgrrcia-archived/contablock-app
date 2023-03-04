package dailyledger

import (
	"context"
	"errors"
	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"

	"github.com/asaskevich/govalidator"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

type DailyLedgerUseCase interface {
	ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error)
}

func NewDefaultDailyLedgerUseCase(options DefaultDailyLedgerUseCaseOptions) (*DefaultDailyLedgerUseCase, error) {
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided AccountUseCase is nil")
	}
	if options.EntryUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided EntryUseCase is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided ContaBlockUseCase is nil")
	}
	if options.ContractUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided ContractUseCase is nil")
	}
	if options.UserUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided UserUseCase is nil")
	}

	return &DefaultDailyLedgerUseCase{
		accountUseCase:    options.AccountUseCase,
		entryUseCase:      options.EntryUseCase,
		contaBlockUseCase: options.ContaBlockUseCase,
		contractUseCase:   options.ContractUseCase,
		userUseCase:       options.UserUseCase,
	}, nil
}

type DefaultDailyLedgerUseCase struct {
	accountUseCase    account.AccountUseCase
	entryUseCase      entry.EntryUseCase
	contaBlockUseCase contablock.ContaBlockUseCase
	contractUseCase   contract.ContractUseCase
	userUseCase       user.UserUseCase
}

type DefaultDailyLedgerUseCaseOptions struct {
	AccountUseCase    account.AccountUseCase
	EntryUseCase      entry.EntryUseCase
	ContaBlockUseCase contablock.ContaBlockUseCase
	ContractUseCase   contract.ContractUseCase
	UserUseCase       user.UserUseCase
}

var _ DailyLedgerUseCase = (*DefaultDailyLedgerUseCase)(nil)

func (useCase *DefaultDailyLedgerUseCase) ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
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

	var listMultipleEntriesIdsInput contablock.ListMultipleEntriesIdsInput
	listMultipleEntriesIdsInput.ContractAddress = storedContract.Contract.EthAddress
	listMultipleEntriesIdsInput.UserAddress = storedUser.User.EthAddress
	listMultipleEntriesIdsInput.UserPrivateKey = storedUser.User.PrivateKey
	listOutput, err := useCase.contaBlockUseCase.ListMultipleEntriesIds(listMultipleEntriesIdsInput)
	if err != nil {
		return nil, err
	}

	var output ListDailyLedgerOutput
	for _, multipleEntryID := range listOutput.MultipleEntriesIds {
		var listEntriesInput entry.ListEntryInput
		listEntriesInput.Project = input.Project
		listEntriesInput.MultipleEntryID = &multipleEntryID
		listEntries, err := useCase.entryUseCase.ListEntries(ctx, listEntriesInput)
		if err != nil {
			return nil, errors.New("failure listing entries")
		}

		if len(listEntries.Items) == 0 {
			continue
		}

		var record Record
		record.ID = multipleEntryID
		for _, entry := range listEntries.Items {
			record.TotalDebitAmount += entry.DebitBalance
			record.TotalCreditAmount += entry.CreditBalance

			var getAccountInput account.DescribeAccountInput
			getAccountInput.ID = entry.AccountID
			getAccountInput.Project = input.Project
			account, err := useCase.accountUseCase.DescribeAccount(ctx, getAccountInput)
			if err != nil {
				return nil, errors.New("failure describing account")
			}

			var entryExpanded EntryExpanded
			entryExpanded.Entry = entry
			entryExpanded.Account = account.Account

			record.EntriesExpanded = append(record.EntriesExpanded, entryExpanded)
		}
		output.Items = append(output.Items, record)
	}

	return &output, err
}
