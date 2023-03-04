package input

import (
	"context"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultGeneralLedgerHttpHandler(options DefaultGeneralLedgerHttpHandlerOptions) (*DefaultGeneralLedgerHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultGeneralLedgerHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultGeneralLedgerHttpHandler struct {
	adapter input.GeneralLedgerInputAdapter
}

type DefaultGeneralLedgerHttpHandlerOptions struct {
	Adapter input.GeneralLedgerInputAdapter
}

func (handler *DefaultGeneralLedgerHttpHandler) RegisterGeneralLedgerRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/projects/{project}/general-ledger", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListGeneralLedgerRequest
		inputPayload.Project = vars["project"]
		inputPayload.ContractID = r.URL.Query().Get("contractId")
		inputPayload.UserID = r.URL.Query().Get("userId")

		outputPayload, err := handler.adapter.AdaptListGeneralLedger(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListGeneralLedgerResponse)
	}).Methods("GET")
}
