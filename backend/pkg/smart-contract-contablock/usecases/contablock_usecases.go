package contablock

import (
	"errors"
	"sync"
)

type ContaBlockUseCase interface {
	GrantAccess(input GrantAccessInput) error
	CreateEntry(input CreateEntryInput) error
	GetEntryData(input GetEntryDataInput) (*GetEntryDataOutput, error)
	ListMultipleEntriesIds(input ListMultipleEntriesIdsInput) (*ListMultipleEntriesIdsOutput, error)
	ListAccountIds(input ListAccountIdsInput) (*ListAccountIdsOutput, error)
	PublishEntry(input PublishEntryInput) error
	RemoveEntry(input RemoveEntryInput) error
	ReviewEntry(input ReviewEntryInput) error
	ApproveEntry(input ApproveEntryInput) error
	RejectEntry(input RejectEntryInput) error
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

var mutex sync.Mutex

var _ ContaBlockUseCase = (*DefaultContaBlockUseCase)(nil)

func (useCase DefaultContaBlockUseCase) GrantAccess(input GrantAccessInput) error {
	var grantAccessAdapterInput GrantAccessAdapterInput
	grantAccessAdapterInput.ContractAddress = input.ContractAddress
	grantAccessAdapterInput.UserPrivateKey = input.UserPrivateKey
	grantAccessAdapterInput.UserAddress = input.UserAddress
	grantAccessAdapterInput.AdminAddress = input.AdminAddress

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.GrantAccess(grantAccessAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) CreateEntry(input CreateEntryInput) error {
	var createEntryAdapterInput CreateEntryAdapterInput
	createEntryAdapterInput.ContractAddress = input.ContractAddress
	createEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	createEntryAdapterInput.UserAddress = input.UserAddress
	createEntryAdapterInput.ID = input.ID
	createEntryAdapterInput.ContractID = input.ContractID
	createEntryAdapterInput.UserID = input.UserID
	createEntryAdapterInput.AccountID = input.AccountID
	createEntryAdapterInput.MultipleEntryID = input.MultipleEntryID
	createEntryAdapterInput.Currency = input.Currency
	createEntryAdapterInput.Concept = input.Concept
	createEntryAdapterInput.Date = input.Date
	createEntryAdapterInput.DebitBalance = input.DebitBalance
	createEntryAdapterInput.CreditBalance = input.CreditBalance

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.CreateEntry(createEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) GetEntryData(input GetEntryDataInput) (*GetEntryDataOutput, error) {
	var getEntryDataAdapterInput GetEntryDataAdapterInput
	getEntryDataAdapterInput.ContractAddress = input.ContractAddress
	getEntryDataAdapterInput.UserPrivateKey = input.UserPrivateKey
	getEntryDataAdapterInput.UserAddress = input.UserAddress
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

func (useCase DefaultContaBlockUseCase) ListMultipleEntriesIds(input ListMultipleEntriesIdsInput) (*ListMultipleEntriesIdsOutput, error) {
	var listMultipleEntriesIdsAdapterInput ListMultipleEntriesIdsAdapterInput
	listMultipleEntriesIdsAdapterInput.ContractAddress = input.ContractAddress
	listMultipleEntriesIdsAdapterInput.UserPrivateKey = input.UserPrivateKey
	listMultipleEntriesIdsAdapterInput.UserAddress = input.UserAddress
	multipleEntriesIds, err := useCase.contaBlockContractAdapter.ListMultipleEntriesIds(listMultipleEntriesIdsAdapterInput)
	if err != nil {
		return nil, err
	}

	var output ListMultipleEntriesIdsOutput
	output.MultipleEntriesIds = append(output.MultipleEntriesIds, multipleEntriesIds.MultipleEntriesIds...)

	return &output, nil
}

func (useCase DefaultContaBlockUseCase) ListAccountIds(input ListAccountIdsInput) (*ListAccountIdsOutput, error) {
	var listAccountIdsAdapterInput ListAccountIdsAdapterInput
	listAccountIdsAdapterInput.ContractAddress = input.ContractAddress
	listAccountIdsAdapterInput.UserPrivateKey = input.UserPrivateKey
	listAccountIdsAdapterInput.UserAddress = input.UserAddress
	accountIds, err := useCase.contaBlockContractAdapter.ListAccountIds(listAccountIdsAdapterInput)
	if err != nil {
		return nil, err
	}

	var output ListAccountIdsOutput
	output.AccountIds = append(output.AccountIds, accountIds.AccountIds...)

	return &output, nil
}

func (useCase DefaultContaBlockUseCase) PublishEntry(input PublishEntryInput) error {
	var publishEntryAdapterInput PublishEntryAdapterInput
	publishEntryAdapterInput.ContractAddress = input.ContractAddress
	publishEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	publishEntryAdapterInput.UserAddress = input.UserAddress
	publishEntryAdapterInput.ID = input.ID

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.PublishEntry(publishEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) RemoveEntry(input RemoveEntryInput) error {
	var removeEntryAdapterInput RemoveEntryAdapterInput
	removeEntryAdapterInput.ContractAddress = input.ContractAddress
	removeEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	removeEntryAdapterInput.UserAddress = input.UserAddress
	removeEntryAdapterInput.ID = input.ID

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.RemoveEntry(removeEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) ReviewEntry(input ReviewEntryInput) error {
	var reviewEntryAdapterInput ReviewEntryAdapterInput
	reviewEntryAdapterInput.ContractAddress = input.ContractAddress
	reviewEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	reviewEntryAdapterInput.UserAddress = input.UserAddress
	reviewEntryAdapterInput.ID = input.ID

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.ReviewEntry(reviewEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) ApproveEntry(input ApproveEntryInput) error {
	var approveEntryAdapterInput ApproveEntryAdapterInput
	approveEntryAdapterInput.ContractAddress = input.ContractAddress
	approveEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	approveEntryAdapterInput.UserAddress = input.UserAddress
	approveEntryAdapterInput.ID = input.ID

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.ApproveEntry(approveEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}

func (useCase DefaultContaBlockUseCase) RejectEntry(input RejectEntryInput) error {
	var rejectEntryAdapterInput RejectEntryAdapterInput
	rejectEntryAdapterInput.ContractAddress = input.ContractAddress
	rejectEntryAdapterInput.UserPrivateKey = input.UserPrivateKey
	rejectEntryAdapterInput.UserAddress = input.UserAddress
	rejectEntryAdapterInput.ID = input.ID

	mutex.Lock()
	err := useCase.contaBlockContractAdapter.RejectEntry(rejectEntryAdapterInput)
	if err != nil {
		return err
	}
	mutex.Unlock()

	return nil
}
