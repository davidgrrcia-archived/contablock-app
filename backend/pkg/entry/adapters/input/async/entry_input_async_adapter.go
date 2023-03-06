package entryInputAsyncAdapter

import (
	"context"
	"errors"
	"fmt"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"

	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
)

type EntryInputAsyncAdapter interface {
	AdaptEntryCreatedEvent(ctx context.Context, event contablockClient.ContablockEntryCreated) error
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

func (adapter DefaultAsyncInputAdapter) AdaptEntryCreatedEvent(ctx context.Context, event contablockClient.ContablockEntryCreated) error {
	fmt.Println("event:", event)

	return nil
}
