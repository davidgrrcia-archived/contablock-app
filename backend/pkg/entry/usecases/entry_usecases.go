package entry

import (
	"context"
	"errors"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	"github.com/google/uuid"
	"time"

	"github.com/asaskevich/govalidator"
)

type EntryUseCase interface {
	CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error)
	DescribeEntry(ctx context.Context, input DescribeEntryInput) (*DescribeEntryOutput, error)
	ListEntries(ctx context.Context, input ListEntryInput) (*ListEntryOutput, error)
	UpdateEntry(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error)
	RemoveEntry(ctx context.Context, input RemoveEntryInput) (*RemoveEntryOutput, error)
}

func NewDefaultEntryUseCase(options DefaultEntryUseCaseOptions) (*DefaultEntryUseCase, error) {
	if options.EntryStorage == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided EntryStorage is nil")
	}
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of EntryUseCase because provided AccountUseCase is nil")
	}

	return &DefaultEntryUseCase{
		entryStorage:   options.EntryStorage,
		accountUseCase: options.AccountUseCase,
	}, nil
}

type DefaultEntryUseCase struct {
	entryStorage   EntryStorage
	accountUseCase account.AccountUseCase
}

type DefaultEntryUseCaseOptions struct {
	EntryStorage   EntryStorage
	AccountUseCase account.AccountUseCase
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

	var desiredEntry EntryStorageData
	if input.ID == nil {
		id := uuid.New().String()
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
		debitBalance := 0.00
		input.DebitBalance = &debitBalance
	}
	desiredEntry.DebitBalance = *input.DebitBalance
	if input.CreditBalance == nil {
		creditBalance := 0.00
		input.CreditBalance = &creditBalance
	}
	desiredEntry.CreditBalance = *input.CreditBalance
	totalBalance := *input.DebitBalance - *input.CreditBalance
	desiredEntry.TotalBalance = totalBalance
	desiredEntry.State = EntryStateCreated
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	storedEntry, err := useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var output CreateEntryOutput
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	output.Date = storedEntry.Date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	output.CreationDate = storedEntry.CreationDate
	output.LastUpdate = storedEntry.LastUpdate

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
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	output.Date = storedEntry.Date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	output.CreationDate = storedEntry.CreationDate
	output.LastUpdate = storedEntry.LastUpdate

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

	if input.DebitBalance == nil && input.CreditBalance == nil {
		return nil, errors.New("valid debit or credit balance is required")
	}

	if input.DebitBalance != nil && input.CreditBalance != nil {
		return nil, errors.New("only one of debit or credit balance is required")
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.ID

	storedEntry, err := useCase.DescribeEntry(ctx, describeEntryInput)
	if err != nil {
		return nil, err
	}

	var desiredEntry EntryStorageData
	desiredEntry.ID = input.ID
	desiredEntry.AccountID = input.AccountID
	desiredEntry.Currency = input.Currency
	if input.Concept == nil {
		desiredEntry.Concept = storedEntry.Concept
	} else {
		desiredEntry.Concept = input.Concept
	}
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
	totalBalance := *input.DebitBalance - *input.CreditBalance
	desiredEntry.TotalBalance = totalBalance
	desiredEntry.State = input.State
	desiredEntry.CreationDate = storedEntry.CreationDate
	desiredEntry.LastUpdate = time.Now()

	updatedEntry, err := useCase.entryStorage.Edit(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var output UpdateEntryOutput
	output.ID = updatedEntry.ID
	output.AccountID = updatedEntry.AccountID
	output.MultipleEntryID = updatedEntry.MultipleEntryID
	output.Currency = updatedEntry.Currency
	output.Concept = updatedEntry.Concept
	output.Date = updatedEntry.Date
	output.Quarter = updatedEntry.Quarter
	output.DebitBalance = updatedEntry.DebitBalance
	output.CreditBalance = updatedEntry.CreditBalance
	output.TotalBalance = updatedEntry.TotalBalance
	output.State = updatedEntry.State
	output.StateReason = updatedEntry.StateReason
	output.CreationDate = updatedEntry.CreationDate
	output.LastUpdate = updatedEntry.LastUpdate

	return &output, err
}

func (useCase *DefaultEntryUseCase) RemoveEntry(ctx context.Context, input RemoveEntryInput) (*RemoveEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedEntry, err := useCase.entryStorage.Remove(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	var output RemoveEntryOutput
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	output.Date = storedEntry.Date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	output.CreationDate = storedEntry.CreationDate
	output.LastUpdate = storedEntry.LastUpdate

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
