package business

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
)

type BusinessUseCase interface {
	DescribeBusiness(ctx context.Context, input DescribeBusinessInput) (*DescribeBusinessOutput, error)
	ListBusiness(ctx context.Context, input ListBusinessInput) (*ListBusinessOutput, error)
}

func NewDefaultBusinessUseCase(options DefaultBusinessUseCaseOptions) (*DefaultBusinessUseCase, error) {
	if options.BusinessStorage == nil {
		return nil, errors.New("cannot create an instance of BusinessUseCase because provided BusinessStorage is nil")
	}

	return &DefaultBusinessUseCase{
		businessStorage: options.BusinessStorage,
	}, nil
}

type DefaultBusinessUseCase struct {
	businessStorage BusinessStorage
}

type DefaultBusinessUseCaseOptions struct {
	BusinessStorage BusinessStorage
}

var _ BusinessUseCase = (*DefaultBusinessUseCase)(nil)

func (useCase *DefaultBusinessUseCase) DescribeBusiness(ctx context.Context, input DescribeBusinessInput) (*DescribeBusinessOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedBusiness, err := useCase.businessStorage.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeBusinessOutput
	output.Business.ID = storedBusiness.ID
	output.Business.Name = storedBusiness.Name
	output.Business.SocialDenomination = storedBusiness.SocialDenomination
	output.Business.RegisteredOffice = storedBusiness.RegisteredOffice
	output.Business.Cif = storedBusiness.Cif
	output.Business.Logo = storedBusiness.Logo
	output.Business.EthAddress = storedBusiness.EthAddress

	return &output, err
}

func (useCase *DefaultBusinessUseCase) ListBusiness(ctx context.Context, input ListBusinessInput) (*ListBusinessOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedBusiness, err := useCase.businessStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	var items []Business
	for _, item := range storedBusiness.Items {
		var business Business
		business.ID = item.ID
		business.Name = item.Name
		business.SocialDenomination = item.SocialDenomination
		business.RegisteredOffice = item.RegisteredOffice
		business.Cif = item.Cif
		business.Logo = item.Logo
		business.EthAddress = item.EthAddress

		items = append(items, business)
	}

	return &ListBusinessOutput{
		Items: items,
	}, err
}
