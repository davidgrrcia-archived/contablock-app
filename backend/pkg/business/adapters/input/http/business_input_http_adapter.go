package businessInputHttpAdapter

import (
	"context"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
)

type BusinessInputAdapter interface {
	AdaptDescribeBusiness(ctx context.Context, data DescribeBusinessRequest) (*DescribeBusinessResponseWrapper, error)
	AdaptListBusiness(ctx context.Context, data ListBusinessRequest) (*ListBusinessResponseWrapper, error)
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
	inputData.ID = data.ID

	outputData, err := adapter.businessUseCase.DescribeBusiness(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeBusinessResponse
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

	outputData, err := adapter.businessUseCase.ListBusiness(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListBusinessResponse
	response.Items = make([]DescribeBusinessResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeBusinessResponse
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
