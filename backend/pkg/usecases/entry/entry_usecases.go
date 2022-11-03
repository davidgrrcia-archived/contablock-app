package entry

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

type EntryUseCase interface {
	CreateEntryUseCase(ctx context.Context, input CreateEntryUseCaseInput) (*CreateEntryUseCaseOutput, error)
	DescribeEntryUseCase(ctx context.Context, input DescribeEntryUseCaseInput) (*DescribeEntryUseCaseOutput, error)
	ListEntryUseCase(ctx context.Context, input ListEntryUseCaseInput) (*ListEntryUseCaseOutput, error)
	UpdateEntryUseCase(ctx context.Context, input UpdateEntryUseCaseInput) (*UpdateEntryUseCaseOutput, error)
	RemoveEntryUseCase(ctx context.Context, input RemoveEntryUseCaseInput) (*RemoveEntryUseCaseOutput, error)
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

func (useCase *DefaultEntryUseCase) CreateEntryUseCase(ctx context.Context, input CreateEntryUseCaseInput) (*CreateEntryUseCaseOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var checkEntry GetInput
	checkEntry.ID = input.ID
	_, err = useCase.entryStorage.Get(ctx, checkEntry)
	if err == nil {
		return nil, errors.New("entry already exists")
	}

	var desiredEntry AddInput
	desiredEntry.ID = input.ID
	desiredEntry.Entryname = input.Entryname
	if err != nil {
		return nil, err
	}

	desiredEntry.Email = input.Email
	desiredEntry.CreationDate = time.Now()
	desiredEntry.LastUpdate = desiredEntry.CreationDate

	storedEntry, err := useCase.entryStorage.Add(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &CreateEntryUseCaseOutput{
		Entry: storedEntry.Entry,
	}, nil
}

func (useCase *DefaultEntryUseCase) DescribeEntryUseCase(ctx context.Context, input DescribeEntryUseCaseInput) (*DescribeEntryUseCaseOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var desiredEntry GetInput
	desiredEntry.ID = input.ID

	storedEntry, err := useCase.entryStorage.Get(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &DescribeEntryUseCaseOutput{
		Entry: storedEntry.Entry,
	}, err
}

func (useCase *DefaultEntryUseCase) ListEntryUseCase(ctx context.Context, input ListEntryUseCaseInput) (*ListEntryUseCaseOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var desiredEntry AllInput

	entryList, err := useCase.entryStorage.All(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	var items = entryList.Items

	return &ListEntryUseCaseOutput{
		Items: items,
	}, err
}

func (useCase *DefaultEntryUseCase) UpdateEntryUseCase(ctx context.Context, input UpdateEntryUseCaseInput) (*UpdateEntryUseCaseOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var desiredEntry EditInput
	desiredEntry.ID = input.ID
	desiredEntry.Entryname = input.Entryname
	if err != nil {
		return nil, err
	}
	desiredEntry.Email = input.Email
	desiredEntry.LastUpdate = time.Now()

	storedEntry, err := useCase.entryStorage.Edit(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &UpdateEntryUseCaseOutput{
		Entry: storedEntry.Entry,
	}, err
}

func (useCase *DefaultEntryUseCase) RemoveEntryUseCase(ctx context.Context, input RemoveEntryUseCaseInput) (*RemoveEntryUseCaseOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var desiredEntry RemoveInput
	desiredEntry.ID = input.ID

	storedEntry, err := useCase.entryStorage.Remove(ctx, desiredEntry)
	if err != nil {
		return nil, err
	}

	return &RemoveEntryUseCaseOutput{
		Entry: storedEntry.Entry,
	}, err
}
