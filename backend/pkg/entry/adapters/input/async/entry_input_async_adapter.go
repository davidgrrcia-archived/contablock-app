package entryInputAsyncAdapter

import (
	"context"
	"errors"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils"
)

type EntryInputAsyncAdapter interface {
	AdaptEntryCreatedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryCreated) error
}

func NewDefaultAsyncInputAdapter(options DefaultAsyncInputAdapterOptions) (*DefaultAsyncInputAdapter, error) {
	if options.EntryUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultAsyncInputAdapter{
		entryUseCase: options.EntryUseCase,
	}, nil
}

type DefaultAsyncInputAdapter struct {
	entryUseCase entry.EntryUseCase
}

type DefaultAsyncInputAdapterOptions struct {
	EntryUseCase entry.EntryUseCase
}

var _ EntryInputAsyncAdapter = (*DefaultAsyncInputAdapter)(nil)

func (adapter DefaultAsyncInputAdapter) AdaptEntryCreatedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryCreated) error {
	debitBalance, err := utils.TransformBigIntToFloat64(event.Entry.DebitBalance)
	if err != nil {
		return err
	}
	creditBalance, err := utils.TransformBigIntToFloat64(event.Entry.CreditBalance)
	if err != nil {
		return err
	}

	var entryCreatedEvent entry.EntryCreatedEventInput
	entryCreatedEvent.Project = project
	entryCreatedEvent.ContractID = event.Entry.ContractId
	entryCreatedEvent.UserID = event.Entry.UserId
	entryCreatedEvent.ID = event.Entry.Id
	entryCreatedEvent.AccountID = event.Entry.AccountId
	entryCreatedEvent.MultipleEntryID = event.Entry.MultipleEntryId
	entryCreatedEvent.Currency = event.Entry.Currency
	entryCreatedEvent.Concept = event.Entry.Concept
	entryCreatedEvent.Date = utils.TransformBigIntToDate(event.Entry.Date)
	entryCreatedEvent.DebitBalance = debitBalance
	entryCreatedEvent.CreditBalance = creditBalance

	err = adapter.entryUseCase.EntryCreatedEvent(ctx, entryCreatedEvent)
	if err != nil {
		return err
	}

	return nil
}
