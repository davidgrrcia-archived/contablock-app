package input

import (
	"context"
	"errors"
	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultAccountsHttpHandler(options DefaultAccountsHttpHandlerOptions) (*DefaultAccountsHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultAccountsHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultAccountsHttpHandler struct {
	adapter input.AccountInputAdapter
}

type DefaultAccountsHttpHandlerOptions struct {
	Adapter input.AccountInputAdapter
}

func (handler *DefaultAccountsHttpHandler) RegisterAccountRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/accounts", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		outputPayload, err := handler.adapter.AdaptListAccounts(ctx, input.ListAccountRequest{})
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListAccountResponse)
	}).Methods("GET")

	r.HandleFunc("/api/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeAccountRequest
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptDescribeAccount(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeAccountResponse)
	}).Methods("GET")
}
