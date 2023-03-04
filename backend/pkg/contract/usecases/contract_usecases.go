package contract

import (
	"context"
	"errors"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/google/uuid"
	"time"

	"github.com/asaskevich/govalidator"
)

type ContractUseCase interface {
	DescribeContract(ctx context.Context, input DescribeContractInput) (*DescribeContractOutput, error)
	ListContracts(ctx context.Context, input ListContractInput) (*ListContractOutput, error)
	CreateContract(ctx context.Context, input CreateContractInput) (*CreateContractOutput, error)
}

func NewDefaultContractUseCase(options DefaultContractUseCaseOptions) (*DefaultContractUseCase, error) {
	if options.ContractStorage == nil {
		return nil, errors.New("cannot create an instance of ContractUseCase because provided ContractStorage is nil")
	}
	if options.ContaBlockUseCase == nil {
		return nil, errors.New("cannot create an instance of ContractUseCase because provided ContaBlockUseCase is nil")
	}

	return &DefaultContractUseCase{
		contractStorage: options.ContractStorage,
	}, nil
}

type DefaultContractUseCase struct {
	contractStorage   ContractStorage
	contaBlockUseCase contablock.ContaBlockUseCase
}

type DefaultContractUseCaseOptions struct {
	ContractStorage   ContractStorage
	ContaBlockUseCase contablock.ContaBlockUseCase
}

var _ ContractUseCase = (*DefaultContractUseCase)(nil)

func (useCase *DefaultContractUseCase) DescribeContract(ctx context.Context, input DescribeContractInput) (*DescribeContractOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedContract, err := useCase.contractStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeContractOutput
	output.Contract.Project = storedContract.Project
	output.Contract.ID = storedContract.ID
	output.Contract.Name = storedContract.Name
	output.Contract.Description = storedContract.Description
	output.Contract.EthAddress = storedContract.EthAddress
	output.Contract.Network = storedContract.Network
	output.Contract.AdminID = storedContract.AdminID
	output.Contract.CreationDate = storedContract.CreationDate
	output.Contract.LastUpdate = storedContract.LastUpdate

	return &output, err
}

func (useCase *DefaultContractUseCase) ListContracts(ctx context.Context, input ListContractInput) (*ListContractOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var filters Filters
	storedContracts, err := useCase.contractStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	var contracts []Contract

	for _, item := range storedContracts.Items {
		var contract Contract
		contract.Project = item.Project
		contract.ID = item.ID
		contract.Name = item.Name
		contract.Description = item.Description
		contract.EthAddress = item.EthAddress
		contract.Network = item.Network
		contract.AdminID = item.AdminID
		contract.CreationDate = item.CreationDate
		contract.LastUpdate = item.LastUpdate
		contracts = append(contracts, contract)
	}

	return &ListContractOutput{
		Items: contracts,
	}, err
}

func (useCase *DefaultContractUseCase) CreateContract(ctx context.Context, input CreateContractInput) (*CreateContractOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	//TODO if the project in the path is the admin, then allow to add

	storedContract, err := useCase.contractStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedContract != nil {
		return nil, errors.New("contract [" + storedContract.ID + "] already exists")
	}

	if input.ID == "" {
		input.ID = uuid.New().String()
	}

	var desiredContract ContractStorageData
	desiredContract.Project = input.Project
	desiredContract.ID = input.ID
	desiredContract.Name = input.Name
	desiredContract.Description = input.Description
	desiredContract.EthAddress = input.EthAddress
	desiredContract.Network = input.Network
	desiredContract.AdminID = input.AdminID
	desiredContract.CreationDate = time.Now()
	desiredContract.LastUpdate = desiredContract.CreationDate

	addedContract, err := useCase.contractStorage.Add(ctx, desiredContract)
	if err != nil {
		return nil, err
	}

	var output CreateContractOutput
	output.Contract.Project = addedContract.Project
	output.Contract.ID = addedContract.ID
	output.Contract.Name = addedContract.Name
	output.Contract.Description = addedContract.Description
	output.Contract.EthAddress = addedContract.EthAddress
	output.Contract.Network = addedContract.Network
	output.Contract.AdminID = addedContract.AdminID
	output.Contract.CreationDate = addedContract.CreationDate
	output.Contract.LastUpdate = addedContract.LastUpdate

	return &output, nil
}
