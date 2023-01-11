package contablock

import (
	"errors"
)

type ContaBlockUseCase interface {
	CreateEntry(input CreateEntryInput) error
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
