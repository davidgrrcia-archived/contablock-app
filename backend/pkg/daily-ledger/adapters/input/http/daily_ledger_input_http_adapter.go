package dailyLedgerInputHttpAdapter

import (
	"context"
	"errors"
	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
)

type DailyLedgerInputAdapter interface {
	AdaptListDailyLedgers(ctx context.Context, data ListDailyLedgerRequest) (*ListDailyLedgerResponseWrapper, error)
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

func (adapter *DefaultHttpInputAdapter) AdaptListDailyLedgers(ctx context.Context, data ListDailyLedgerRequest) (*ListDailyLedgerResponseWrapper, error) {
	var inputData dailyLedger.ListDailyLedgerInput

	outputData, err := adapter.dailyLedgerUseCase.ListDailyLedgers(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListDailyLedgerResponse
	response.Items = make([]DescribeDailyLedgerResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeDailyLedgerResponse
		responseItem.ID = item.ID
		responseItem.Name = item.Name
		responseItem.Level = item.Level
		responseItem.Type = item.Type
		responseItem.DisplayName = item.DisplayName
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListDailyLedgerResponseWrapper{}
	responseWrapper.ListDailyLedgerResponse = response
	return &responseWrapper, nil
}
