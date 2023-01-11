package contablock

import (
	"errors"
)

type ContaBlockUseCase interface {
	CreateEntry(input CreateEntryInput) error
	GetEntryData(input GetEntryDataInput) (*GetEntryDataOutput, error)
	ListMultipleEntriesIds() (*ListMultipleEntriesIdsOutput, error)
	ListAccountIds() (*ListAccountIdsOutput, error)
	PublishEntry(input PublishEntryInput) error
}

func NewDefaultContaBlockUseCase(options DefaultContaBlockUseCaseOptions) (*DefaultContaBlockUseCase, error) {
	if options.ContaBlockContractAdapter == nil {
		return nil, errors.New("cannot create an instance of ContaBlockUseCase because provided ContaBlockContractAdapter is nil")
	}

	return &DefaultContaBlockUseCase{
		contaBlockContractAdapter: options.ContaBlockContractAdapter,
	}, nil
}

type DefaultContaBlockUseCase struct {
	contaBlockContractAdapter ContaBlockContractAdapter
}

type DefaultContaBlockUseCaseOptions struct {
	ContaBlockContractAdapter ContaBlockContractAdapter
}

var _ ContaBlockUseCase = (*DefaultContaBlockUseCase)(nil)

func (useCase DefaultContaBlockUseCase) CreateEntry(input CreateEntryInput) error {
	var createEntryAdapterInput CreateEntryAdapterInput
	createEntryAdapterInput.ID = input.ID
	createEntryAdapterInput.AccountID = input.AccountID
	createEntryAdapterInput.MultipleEntryID = input.MultipleEntryID
	createEntryAdapterInput.Currency = input.Currency
	createEntryAdapterInput.Concept = input.Concept
	createEntryAdapterInput.Date = input.Date
	createEntryAdapterInput.DebitBalance = input.DebitBalance
	createEntryAdapterInput.CreditBalance = input.CreditBalance

	err := useCase.contaBlockContractAdapter.CreateEntry(createEntryAdapterInput)
	if err != nil {
		return err
	}

	return nil
}

func (useCase DefaultContaBlockUseCase) GetEntryData(input GetEntryDataInput) (*GetEntryDataOutput, error) {
	var getEntryDataAdapterInput GetEntryDataAdapterInput
	getEntryDataAdapterInput.ID = input.ID

	entryData, err := useCase.contaBlockContractAdapter.GetEntryData(getEntryDataAdapterInput)
	if err != nil {
		return nil, err
	}

	var output GetEntryDataOutput
	output.ID = entryData.ID
	output.AccountID = entryData.AccountID
	output.MultipleEntryID = entryData.MultipleEntryID
	output.Currency = entryData.Currency
	output.Concept = entryData.Concept
	output.Date = entryData.Date
	output.DebitBalance = entryData.DebitBalance
	output.CreditBalance = entryData.CreditBalance
	output.State = entryData.State

	return &output, nil
}

func (useCase DefaultContaBlockUseCase) ListMultipleEntriesIds() (*ListMultipleEntriesIdsOutput, error) {
	multipleEntriesIds, err := useCase.contaBlockContractAdapter.ListMultipleEntriesIds()
	if err != nil {
		return nil, err
	}

	var output ListMultipleEntriesIdsOutput
	for _, item := range multipleEntriesIds.MultipleEntriesIds {
		output.multipleEntriesIds = append(output.multipleEntriesIds, item)
	}

	return &output, nil
}

func (useCase DefaultContaBlockUseCase) ListAccountIds() (*ListAccountIdsOutput, error) {
	accountIds, err := useCase.contaBlockContractAdapter.ListAccountIds()
	if err != nil {
		return nil, err
	}

	var output ListAccountIdsOutput
	for _, item := range accountIds.AccountIds {
		output.accountIds = append(output.accountIds, item)
	}

	return &output, nil
}

func (useCase DefaultContaBlockUseCase) PublishEntry(input PublishEntryInput) error {
	var publishEntryAdapterInput PublishEntryAdapterInput
	publishEntryAdapterInput.ID = input.ID

	err := useCase.contaBlockContractAdapter.PublishEntry(publishEntryAdapterInput)
	if err != nil {
		return err
	}

	return nil
}
