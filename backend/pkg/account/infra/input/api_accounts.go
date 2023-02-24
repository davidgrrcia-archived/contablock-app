package input

import (
	"context"
	"encoding/json"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

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
	r.HandleFunc("/api/projects/{project}/accounts", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListAccountRequest
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptListAccounts(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListAccountResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeAccountRequest
		inputPayload.ID = vars["id"]
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptDescribeAccount(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeAccountResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/accounts", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateAccountRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateAccount(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateAccountResponse)
	}).Methods("POST")
}
