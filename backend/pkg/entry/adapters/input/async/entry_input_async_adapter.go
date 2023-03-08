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
	AdaptEntryRemovedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryRemoved) error
	AdaptEntryReviewRequestedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryReviewRequested) error
	AdaptEntryApprovedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryApproved) error
	AdaptEntryRejectedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryRejected) error
	AdaptEntryPublishedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryPublished) error
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

func (adapter DefaultAsyncInputAdapter) AdaptEntryRemovedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryRemoved) error {
	var entryRemovedEvent entry.EntryRemovedEventInput
	entryRemovedEvent.Project = project
	entryRemovedEvent.ID = event.Id

	err := adapter.entryUseCase.EntryRemovedEvent(ctx, entryRemovedEvent)
	if err != nil {
		return err
	}

	return nil
}

func (adapter DefaultAsyncInputAdapter) AdaptEntryReviewRequestedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryReviewRequested) error {
	var entryReviewRequestedEvent entry.EntryReviewRequestedEventInput
	entryReviewRequestedEvent.Project = project
	entryReviewRequestedEvent.ID = event.Id

	err := adapter.entryUseCase.EntryReviewRequestedEvent(ctx, entryReviewRequestedEvent)
	if err != nil {
		return err
	}

	return nil
}

func (adapter DefaultAsyncInputAdapter) AdaptEntryApprovedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryApproved) error {
	var entryApprovedEvent entry.EntryApprovedEventInput
	entryApprovedEvent.Project = project
	entryApprovedEvent.ID = event.Id

	err := adapter.entryUseCase.EntryApprovedEvent(ctx, entryApprovedEvent)
	if err != nil {
		return err
	}

	return nil
}

func (adapter DefaultAsyncInputAdapter) AdaptEntryRejectedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryRejected) error {
	var entryRejectedEvent entry.EntryRejectedEventInput
	entryRejectedEvent.Project = project
	entryRejectedEvent.ID = event.Id

	err := adapter.entryUseCase.EntryRejectedEvent(ctx, entryRejectedEvent)
	if err != nil {
		return err
	}

	return nil
}

func (adapter DefaultAsyncInputAdapter) AdaptEntryPublishedEvent(ctx context.Context, project string, event contablockClient.ContablockEntryPublished) error {
	var entryPublishedEvent entry.EntryPublishedEventInput
	entryPublishedEvent.Project = project
	entryPublishedEvent.ID = event.Id

	err := adapter.entryUseCase.EntryPublishedEvent(ctx, entryPublishedEvent)
	if err != nil {
		return err
	}

	return nil
}
