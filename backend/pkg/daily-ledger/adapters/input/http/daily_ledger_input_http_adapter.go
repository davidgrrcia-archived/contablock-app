package dailyLedgerInputHttpAdapter

import (
	"context"
	"errors"

	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
)

type DailyLedgerInputAdapter interface {
	AdaptListDailyLedger(ctx context.Context, data ListDailyLedgerRequest) (*ListDailyLedgerResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.DailyLedgerUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		dailyLedgerUseCase: options.DailyLedgerUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	dailyLedgerUseCase dailyledger.DailyLedgerUseCase
}

type DefaultHttpInputAdapterOptions struct {
	DailyLedgerUseCase dailyledger.DailyLedgerUseCase
}

var _ DailyLedgerInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptListDailyLedger(ctx context.Context, data ListDailyLedgerRequest) (*ListDailyLedgerResponseWrapper, error) {
	var inputData dailyledger.ListDailyLedgerInput
	outputData, err := adapter.dailyLedgerUseCase.ListDailyLedger(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListDailyLedgerResponse
	response.Items = make([]Record, 0)

	for _, item := range outputData.Items {
		var responseItem Record
		responseItem.Project = item.Project
		responseItem.ID = item.ID

		for _, entry := range item.EntriesExpanded {
			var entryExpanded EntryExpanded
			entryExpanded.Entry.Project = entry.Entry.Project
			entryExpanded.Entry.ID = entry.Entry.ID
			entryExpanded.Entry.AccountID = entry.Entry.AccountID
			entryExpanded.Entry.MultipleEntryID = entry.Entry.MultipleEntryID
			entryExpanded.Entry.Currency = entry.Entry.Currency
			entryExpanded.Entry.Concept = entry.Entry.Concept
			entryExpanded.Entry.Date = entry.Entry.Date.Unix()
			entryExpanded.Entry.Quarter = entry.Entry.Quarter
			entryExpanded.Entry.DebitBalance = entry.Entry.DebitBalance
			entryExpanded.Entry.CreditBalance = entry.Entry.CreditBalance
			entryExpanded.Entry.TotalBalance = entry.Entry.TotalBalance
			entryExpanded.Entry.State = entry.Entry.State
			entryExpanded.Entry.StateReason = entry.Entry.StateReason
			entryExpanded.Entry.CreationDate = entry.Entry.CreationDate.Unix()
			entryExpanded.Entry.LastUpdate = entry.Entry.LastUpdate.Unix()
			entryExpanded.Account.Project = entry.Account.Project
			entryExpanded.Account.ID = entry.Account.ID
			entryExpanded.Account.Name = entry.Account.Name
			entryExpanded.Account.Level = entry.Account.Level
			entryExpanded.Account.BalanceID = entry.Account.BalanceID
			entryExpanded.Account.DisplayName = entry.Account.DisplayName

			responseItem.EntriesExpanded = append(responseItem.EntriesExpanded, entryExpanded)
		}

		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListDailyLedgerResponseWrapper{}
	responseWrapper.ListDailyLedgerResponse = response
	return &responseWrapper, nil
}
