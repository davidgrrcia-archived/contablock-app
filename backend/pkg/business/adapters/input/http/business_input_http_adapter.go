package businessInputHttpAdapter

import (
	"context"
	"errors"

	business "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
)

type BusinessInputAdapter interface {
	AdaptDescribeBusiness(ctx context.Context, data DescribeBusinessRequest) (*DescribeBusinessResponseWrapper, error)
	AdaptListBusiness(ctx context.Context, data ListBusinessRequest) (*ListBusinessResponseWrapper, error)
	AdaptCreateBusiness(ctx context.Context, data CreateBusinessRequest) (*CreateBusinessResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.BusinessUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		businessUseCase: options.BusinessUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	businessUseCase business.BusinessUseCase
}

type DefaultHttpInputAdapterOptions struct {
	BusinessUseCase business.BusinessUseCase
}

var _ BusinessInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptDescribeBusiness(ctx context.Context, data DescribeBusinessRequest) (*DescribeBusinessResponseWrapper, error) {
	var inputData business.DescribeBusinessInput
	inputData.Project = data.Project
	inputData.ID = data.ID

	outputData, err := adapter.businessUseCase.DescribeBusiness(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeBusinessResponse
	response.Project = outputData.Business.Project
	response.ID = outputData.Business.ID
	response.Name = outputData.Business.Name
	response.SocialDenomination = outputData.Business.SocialDenomination
	response.RegisteredOffice = outputData.Business.RegisteredOffice
	response.Cif = outputData.Business.Cif
	response.Logo = outputData.Business.Logo
	response.EthAddress = outputData.Business.EthAddress
	responseWrapper := DescribeBusinessResponseWrapper{}
	responseWrapper.DescribeBusinessResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListBusiness(ctx context.Context, data ListBusinessRequest) (*ListBusinessResponseWrapper, error) {
	var inputData business.ListBusinessInput
	inputData.Project = data.Project

	outputData, err := adapter.businessUseCase.ListBusiness(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListBusinessResponse
	response.Items = make([]DescribeBusinessResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeBusinessResponse
		responseItem.Project = item.Project
		responseItem.ID = item.ID
		responseItem.Name = item.Name
		responseItem.SocialDenomination = item.SocialDenomination
		responseItem.RegisteredOffice = item.RegisteredOffice
		responseItem.Cif = item.Cif
		responseItem.Logo = item.Logo
		responseItem.EthAddress = item.EthAddress
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListBusinessResponseWrapper{}
	responseWrapper.ListBusinessResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptCreateBusiness(ctx context.Context, data CreateBusinessRequest) (*CreateBusinessResponseWrapper, error) {
	var inputData business.CreateBusinessInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.Name = data.Name
	inputData.SocialDenomination = data.SocialDenomination
	inputData.RegisteredOffice = data.RegisteredOffice
	inputData.Cif = data.Cif
	inputData.Logo = data.Logo
	inputData.EthAddress = data.EthAddress

	outputData, err := adapter.businessUseCase.CreateBusiness(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateBusinessResponse
	response.Project = outputData.Business.Project
	response.ID = outputData.Business.ID
	response.Name = outputData.Business.Name
	response.SocialDenomination = outputData.Business.SocialDenomination
	response.RegisteredOffice = outputData.Business.RegisteredOffice
	response.Cif = outputData.Business.Cif
	response.Logo = outputData.Business.Logo
	response.EthAddress = outputData.Business.EthAddress

	responseWrapper := CreateBusinessResponseWrapper{}
	responseWrapper.CreateBusinessResponse = response

	return &responseWrapper, nil
}
