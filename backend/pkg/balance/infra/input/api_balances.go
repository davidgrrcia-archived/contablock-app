package input

import (
	"context"
	"errors"
	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultBalancesHttpHandler(options DefaultBalancesHttpHandlerOptions) (*DefaultBalancesHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultBalancesHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultBalancesHttpHandler struct {
	adapter input.BalanceInputAdapter
}

type DefaultBalancesHttpHandlerOptions struct {
	Adapter input.BalanceInputAdapter
}

func (handler *DefaultBalancesHttpHandler) RegisterBalanceRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/balances", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		var inputPayload input.ListBalanceRequest
		inputPayload.Year = r.URL.Query().Get("year")

		outputPayload, err := handler.adapter.AdaptListBalances(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListBalanceResponse)
	}).Methods("GET")
}
