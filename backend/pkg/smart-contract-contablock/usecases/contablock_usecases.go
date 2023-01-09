package contablock

import (
	"errors"
)

type ContaBlockUseCase interface {
	Store(input StoreInput) error
	Retrieve() (*RetrieveOutput, error)
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

func (useCase DefaultContaBlockUseCase) Store(input StoreInput) error {
	var storeAdapterInput StoreAdapterInput
	storeAdapterInput.Num = input.Num

	err := useCase.contaBlockContractAdapter.Store(storeAdapterInput)
	if err != nil {
		return err
	}

	return nil
}

func (useCase DefaultContaBlockUseCase) Retrieve() (*RetrieveOutput, error) {
	value, err := useCase.contaBlockContractAdapter.Retrieve()
	if err != nil {
		return nil, err
	}

	var output RetrieveOutput
	output.Num = value.Num

	return &output, nil
}
