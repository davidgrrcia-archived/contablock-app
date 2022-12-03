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

	return &DescribeBusinessOutput{
		Business: storedBusiness.Business,
	}, err
}

func (useCase *DefaultBusinessUseCase) ListBusiness(ctx context.Context, input ListBusinessInput) (*ListBusinessOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	businessCollection, err := useCase.businessStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	return &ListBusinessOutput{
		Items: businessCollection.Items,
	}, err
}
