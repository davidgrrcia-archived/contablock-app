package business

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
)

type BusinessUseCase interface {
	DescribeBusiness(ctx context.Context, input DescribeBusinessInput) (*DescribeBusinessOutput, error)
	ListBusiness(ctx context.Context, input ListBusinessInput) (*ListBusinessOutput, error)
	CreateBusiness(ctx context.Context, input CreateBusinessInput) (*CreateBusinessOutput, error)
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
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedBusiness, err := useCase.businessStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeBusinessOutput
	output.Business.Project = storedBusiness.Project
	output.Business.ID = storedBusiness.ID
	output.Business.Name = storedBusiness.Name
	output.Business.SocialDenomination = storedBusiness.SocialDenomination
	output.Business.RegisteredOffice = storedBusiness.RegisteredOffice
	output.Business.Cif = storedBusiness.Cif
	output.Business.Logo = storedBusiness.Logo

	return &output, err
}

func (useCase *DefaultBusinessUseCase) ListBusiness(ctx context.Context, input ListBusinessInput) (*ListBusinessOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedBusiness, err := useCase.businessStorage.All(ctx, input.Project)
	if err != nil {
		return nil, err
	}

	var items []Business
	for _, item := range storedBusiness.Items {
		var business Business
		business.Project = item.Project
		business.ID = item.ID
		business.Name = item.Name
		business.SocialDenomination = item.SocialDenomination
		business.RegisteredOffice = item.RegisteredOffice
		business.Cif = item.Cif
		business.Logo = item.Logo

		items = append(items, business)
	}

	return &ListBusinessOutput{
		Items: items,
	}, err
}

func (useCase *DefaultBusinessUseCase) CreateBusiness(ctx context.Context, input CreateBusinessInput) (*CreateBusinessOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	//TODO if the project in the path is the admin, then allow to add

	storedBusiness, err := useCase.businessStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedBusiness != nil {
		return nil, errors.New("business [" + storedBusiness.ID + "] already exists")
	}

	var desiredBusiness BusinessStorageData
	desiredBusiness.Project = input.Project
	desiredBusiness.ID = input.ID
	desiredBusiness.Name = input.Name
	desiredBusiness.SocialDenomination = input.SocialDenomination
	desiredBusiness.RegisteredOffice = input.RegisteredOffice
	desiredBusiness.Cif = input.Cif
	desiredBusiness.Logo = input.Logo

	addedBusiness, err := useCase.businessStorage.Add(ctx, desiredBusiness)
	if err != nil {
		return nil, err
	}

	var output CreateBusinessOutput
	output.Business.Project = addedBusiness.Project
	output.Business.ID = addedBusiness.ID
	output.Business.Name = addedBusiness.Name
	output.Business.SocialDenomination = addedBusiness.SocialDenomination
	output.Business.RegisteredOffice = addedBusiness.RegisteredOffice
	output.Business.Cif = addedBusiness.Cif
	output.Business.Logo = addedBusiness.Logo

	return &output, nil
}
