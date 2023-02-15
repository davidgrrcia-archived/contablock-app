package generalLedgerInputHttpAdapter

import (
	"context"
	"errors"

	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	generalledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/usecases"
)

type GeneralLedgerInputAdapter interface {
	AdaptListGeneralLedger(ctx context.Context, data ListGeneralLedgerRequest) (*ListGeneralLedgerResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.GeneralLedgerUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		generalLedgerUseCase: options.GeneralLedgerUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	generalLedgerUseCase generalledger.GeneralLedgerUseCase
}

type DefaultHttpInputAdapterOptions struct {
	GeneralLedgerUseCase generalledger.GeneralLedgerUseCase
}

var _ GeneralLedgerInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptListGeneralLedger(ctx context.Context, data ListGeneralLedgerRequest) (*ListGeneralLedgerResponseWrapper, error) {
	var inputData generalledger.ListGeneralLedgerInput
	outputData, err := adapter.generalLedgerUseCase.ListGeneralLedger(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListGeneralLedgerResponse
	response.Items = make([]Record, 0)

	for _, item := range outputData.Items {
		var responseItem Record
		responseItem.Account.ID = item.Account.ID
		responseItem.Account.Name = item.Account.Name
		responseItem.Account.BalanceID = item.Account.BalanceID
		responseItem.Account.Level = item.Account.Level
		responseItem.Account.DisplayName = item.Account.DisplayName

		for _, item2 := range item.Entries {
			var entry entry.Entry
			entry.ID = item2.ID
			entry.AccountID = item2.AccountID
			entry.MultipleEntryID = item2.MultipleEntryID
			entry.Currency = item2.Currency
			entry.Concept = item2.Concept
			entry.Date = item2.Date.Unix()
			entry.Quarter = item2.Quarter
			entry.DebitBalance = item2.DebitBalance
			entry.CreditBalance = item2.CreditBalance
			entry.TotalBalance = item2.TotalBalance
			entry.State = item2.State
			entry.StateReason = item2.StateReason
			entry.CreationDate = item2.CreationDate.Unix()
			entry.LastUpdate = item2.LastUpdate.Unix()

			responseItem.Entries = append(responseItem.Entries, entry)
		}

		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListGeneralLedgerResponseWrapper{}
	responseWrapper.ListGeneralLedgerResponse = response
	return &responseWrapper, nil
}
