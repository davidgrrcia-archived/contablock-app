package entry

import (
	"context"
	"errors"
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

	return &DefaultEntryUseCase{
		entryStorage: options.EntryStorage,
	}, nil
}

type DefaultEntryUseCase struct {
	entryStorage EntryStorage
}

type DefaultEntryUseCaseOptions struct {
	EntryStorage EntryStorage
}

var _ EntryUseCase = (*DefaultEntryUseCase)(nil)

func (useCase *DefaultEntryUseCase) CreateEntry(ctx context.Context, input CreateEntryInput) (*CreateEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var describeEntryInput DescribeEntryInput
	describeEntryInput.ID = input.ID
	_, err = useCase.DescribeEntry(ctx, describeEntryInput)
	if err == nil {
		return nil, errors.New("entry already exists")
	}

	var desiredEntry EntryStorageData
	desiredEntry.ID = input.ID
	desiredEntry.AccountID = input.AccountID
	desiredEntry.MultipleEntryID = input.MultipleEntryID
	desiredEntry.Currency = input.Currency
	desiredEntry.Concept = input.Concept
	desiredEntry.Date = input.Date
	quarter := useCase.calculateQuarter(input.Date)
	desiredEntry.Quarter = quarter
	desiredEntry.DebitBalance = input.DebitBalance
	desiredEntry.CreditBalance = input.CreditBalance
	totalBalance := input.DebitBalance - input.CreditBalance
	desiredEntry.TotalBalance = totalBalance
	desiredEntry.State = EntryStateCreated
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	storedEntry, err := useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &CreateEntryOutput{
		Entry: storedEntry.Entry,
	}, nil
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

	return &DescribeEntryOutput{
		Entry: storedEntry.Entry,
	}, err
}

func (useCase *DefaultEntryUseCase) ListEntries(ctx context.Context, input ListEntryInput) (*ListEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	entryCollection, err := useCase.entryStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	return &ListEntryOutput{
		Items: entryCollection.Items,
	}, err
}

func (useCase *DefaultEntryUseCase) UpdateEntry(ctx context.Context, input UpdateEntryInput) (*UpdateEntryOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var desiredEntry EntryStorageData
	desiredEntry.ID = input.ID
	if input.AccountID != nil {
		desiredEntry.AccountID = *input.AccountID
	}
	if input.Currency != nil {
		desiredEntry.Currency = *input.Currency
	}
	if input.Concept != nil {
		desiredEntry.Concept = *input.Concept
	}
	if input.Date != nil {
		desiredEntry.Date = *input.Date
		quarter := useCase.calculateQuarter(*input.Date)
		desiredEntry.Quarter = quarter
	}
	if input.DebitBalance != nil {
		desiredEntry.DebitBalance = *input.DebitBalance
	}
	if input.CreditBalance != nil {
		desiredEntry.CreditBalance = *input.CreditBalance
	}
	if input.DebitBalance != nil && input.CreditBalance != nil {
		totalBalance := *input.DebitBalance - *input.CreditBalance
		desiredEntry.TotalBalance = totalBalance
	}
	if input.State != nil {
		desiredEntry.State = *input.State
	}
	desiredEntry.LastUpdate = time.Now()

	storedEntry, err := useCase.entryStorage.Edit(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &UpdateEntryOutput{
		Entry: storedEntry.Entry,
	}, err
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

	return &RemoveEntryOutput{
		Entry: storedEntry.Entry,
	}, err
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
