package contractInputHttpAdapter

import (
	"context"
	"errors"

	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
)

type ContractInputAdapter interface {
	AdaptCreateContract(ctx context.Context, data CreateContractRequest) (*CreateContractResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.ContractUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		contractUseCase: options.ContractUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	contractUseCase contract.ContractUseCase
}

type DefaultHttpInputAdapterOptions struct {
	ContractUseCase contract.ContractUseCase
}

var _ ContractInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptCreateContract(ctx context.Context, data CreateContractRequest) (*CreateContractResponseWrapper, error) {
	var inputData contract.CreateContractInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.Name = data.Name
	inputData.Description = data.Description
	inputData.EthAddress = data.EthAddress
	inputData.Network = data.Network
	inputData.AdminID = data.AdminID

	outputData, err := adapter.contractUseCase.CreateContract(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateContractResponse
	response.Project = outputData.Contract.Project
	response.ID = outputData.Contract.ID
	response.Name = outputData.Contract.Name
	response.Description = outputData.Contract.Description
	response.EthAddress = outputData.Contract.EthAddress
	response.Network = outputData.Contract.Network
	response.AdminID = outputData.Contract.AdminID
	response.CreationDate = outputData.Contract.CreationDate.Unix()
	response.LastUpdate = outputData.Contract.LastUpdate.Unix()

	responseWrapper := CreateContractResponseWrapper{}
	responseWrapper.CreateContractResponse = response

	return &responseWrapper, nil
}
