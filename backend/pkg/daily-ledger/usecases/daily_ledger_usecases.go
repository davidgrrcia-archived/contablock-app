package dailyledger

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type DailyLedgerUseCase interface {
	ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error)
}

func NewDefaultDailyLedgerUseCase(options DefaultDailyLedgerUseCaseOptions) (*DefaultDailyLedgerUseCase, error) {
	if options.AccountUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided AccountUseCase is nil")
	}
	if options.EntryUseCase == nil {
		return nil, errors.New("cannot create an instance of DailyLedgerUseCase because provided EntryUseCase is nil")
	}

	return &DefaultDailyLedgerUseCase{
		accountUseCase: options.AccountUseCase,
		entryUseCase:   options.EntryUseCase,
	}, nil
}

type DefaultDailyLedgerUseCase struct {
	accountUseCase account.AccountUseCase
	entryUseCase   entry.EntryUseCase
}

type DefaultDailyLedgerUseCaseOptions struct {
	AccountUseCase account.AccountUseCase
	EntryUseCase   entry.EntryUseCase
}

var _ DailyLedgerUseCase = (*DefaultDailyLedgerUseCase)(nil)

func (useCase *DefaultDailyLedgerUseCase) ListDailyLedger(ctx context.Context, input ListDailyLedgerInput) (*ListDailyLedgerOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//todo: implement

	return &ListDailyLedgerOutput{
		Items: items,
	}, err
}
