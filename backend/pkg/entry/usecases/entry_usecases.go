package entry

import (
	"context"
	"errors"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/google/uuid"
	"math/big"
	"time"

	"github.com/asaskevich/govalidator"
)

type EntryUseCase interface {
	CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error)
	DescribeEntry(ctx context.Context, input DescribeEntryInput) (*DescribeEntryOutput, error)
	ListEntries(ctx context.Context, input ListEntryInput) (*ListEntryOutput, error)
	UpdateEntry(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error)
	RemoveEntry(ctx context.Context, input RemoveEntryInput) (*RemoveEntryOutput, error)
	PublishEntry(ctx context.Context, input PublishEntryInput) (*PublishEntryOutput, error)
}

func NewDefaultEntryUseCase(options DefaultEntryUseCaseOptions) (*DefaultEntryUseCase, error) {
	if options.EntryStorage == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided EntryStorage is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided ContaBlockUseCase is nil")
	}

	return &DefaultEntryUseCase{
		entryStorage:      options.EntryStorage,
		contaBlockUseCase: options.ContaBlockUseCase,
	}, nil
}

type DefaultEntryUseCase struct {
	entryStorage      EntryStorage
	contaBlockUseCase contablock.ContaBlockUseCase
}

type DefaultEntryUseCaseOptions struct {
	EntryStorage      EntryStorage
	ContaBlockUseCase contablock.ContaBlockUseCase
}

var _ EntryUseCase = (*DefaultEntryUseCase)(nil)

func (useCase *DefaultEntryUseCase) CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	if input.DebitBalance == nil && input.CreditBalance == nil {
		return nil, errors.New("valid debit or credit balance is required")
	}

	if input.DebitBalance != nil && input.CreditBalance != nil {
		return nil, errors.New("only one of debit or credit balance is required")
	}

	id := uuid.New().String()

	var createEntryInput contablock.CreateEntryInput
	createEntryInput.ID = id
	createEntryInput.AccountID = input.AccountID
	createEntryInput.MultipleEntryID = input.MultipleEntryID
	createEntryInput.Currency = input.Currency
	createEntryInput.Concept = *input.Concept
	createEntryInput.Date = input.Date.String()
	createEntryInput.DebitBalance = input.DebitBalance
	createEntryInput.CreditBalance = input.CreditBalance
	err = useCase.contaBlockUseCase.CreateEntry(createEntryInput)
	if err != nil {
		return nil, err
	}

	var desiredEntry EntryStorageData
	if input.ID == nil {
		input.ID = &id
	}
	desiredEntry.ID = *input.ID
	desiredEntry.AccountID = input.AccountID
	desiredEntry.MultipleEntryID = input.MultipleEntryID
	desiredEntry.Currency = input.Currency
	desiredEntry.Concept = input.Concept
	desiredEntry.Date = input.Date
	quarter := useCase.calculateQuarter(input.Date)
	desiredEntry.Quarter = quarter
	if input.DebitBalance == nil {
		debitBalance := new(big.Int).SetUint64(0.00)
		input.DebitBalance = debitBalance
	}
	desiredEntry.DebitBalance = *input.DebitBalance
	if input.CreditBalance == nil {
		creditBalance := new(big.Int).SetUint64(0.00)
		input.CreditBalance = creditBalance
	}
	desiredEntry.CreditBalance = *input.CreditBalance
	desiredEntry.TotalBalance.Sub(input.DebitBalance, input.CreditBalance)
	desiredEntry.State = EntryStateCreated
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	storedEntry, err := useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var output CreateEntryOutput
	output.Entry.ID = storedEntry.ID
	output.Entry.AccountID = storedEntry.AccountID
	output.Entry.MultipleEntryID = storedEntry.MultipleEntryID
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
		return nil, errors.New("invalid input")
	}

	storedEntry, err := useCase.entryStorage.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeEntryOutput
	output.Entry.ID = storedEntry.ID
	output.Entry.AccountID = storedEntry.AccountID
	output.Entry.MultipleEntryID = storedEntry.MultipleEntryID
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
		return nil, errors.New("invalid input")
	}

	storedEntries, err := useCase.entryStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	var items []Entry
	for _, item := range storedEntries.Items {
		var entry Entry
		entry.ID = item.ID
		entry.AccountID = item.AccountID
		entry.MultipleEntryID = item.MultipleEntryID
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
		return nil, errors.New("invalid input")
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.Entry.ID

	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	var desiredEntry EntryStorageData
	desiredEntry.ID = input.Entry.ID
	desiredEntry.AccountID = input.Entry.AccountID
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
	desiredEntry.TotalBalance.Sub(&input.Entry.DebitBalance, &input.Entry.CreditBalance)
	desiredEntry.State = input.Entry.State
	desiredEntry.CreationDate = storedEntry.Entry.CreationDate
	desiredEntry.LastUpdate = time.Now()

	updatedEntry, err := useCase.entryStorage.Edit(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var output UpdateEntryOutput
	output.Entry.ID = updatedEntry.ID
	output.Entry.AccountID = updatedEntry.AccountID
	output.Entry.MultipleEntryID = updatedEntry.MultipleEntryID
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
		return nil, errors.New("invalid input")
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.ID
	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	var getEntryInput contablock.GetEntryDataInput
	getEntryInput.ID = input.ID
	entryData, err := useCase.contaBlockUseCase.GetEntryData(getEntryInput)
	if err != nil {
		return nil, err
	}

	if !(storedEntry.Entry.State == EntryStateCreated && entryData.State == contablock.EntryStateCreated) {
		return nil, errors.New("failure publishing state not allowed")
	}

	var removeEntryInput contablock.RemoveEntryInput
	removeEntryInput.ID = input.ID
	err = useCase.contaBlockUseCase.RemoveEntry(removeEntryInput)
	if err != nil {
		return nil, err
	}

	removedEntry, err := useCase.entryStorage.Remove(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	var output RemoveEntryOutput
	output.Entry.ID = removedEntry.ID
	output.Entry.AccountID = removedEntry.AccountID
	output.Entry.MultipleEntryID = removedEntry.MultipleEntryID
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
		return nil, errors.New("invalid input")
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.ID
	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	var getEntryInput contablock.GetEntryDataInput
	getEntryInput.ID = input.ID
	entryData, err := useCase.contaBlockUseCase.GetEntryData(getEntryInput)
	if err != nil {
		return nil, err
	}

	if !(storedEntry.Entry.State == EntryStateCreated && entryData.State == contablock.EntryStateCreated) {
		return nil, errors.New("failure publishing state not allowed")
	}

	var publishEntryInput contablock.PublishEntryInput
	publishEntryInput.ID = input.ID
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
