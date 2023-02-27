package input

import (
	"context"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultDailyLedgerHttpHandler(options DefaultDailyLedgerHttpHandlerOptions) (*DefaultDailyLedgerHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultDailyLedgerHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultDailyLedgerHttpHandler struct {
	adapter input.DailyLedgerInputAdapter
}

type DefaultDailyLedgerHttpHandlerOptions struct {
	Adapter input.DailyLedgerInputAdapter
}

func (handler *DefaultDailyLedgerHttpHandler) RegisterDailyLedgerRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/projects/{project}/daily-ledger", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListDailyLedgerRequest
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptListDailyLedger(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListDailyLedgerResponse)
	}).Methods("GET")
}
