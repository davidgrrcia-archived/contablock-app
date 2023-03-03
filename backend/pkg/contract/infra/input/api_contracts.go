package input

import (
	"context"
	"encoding/json"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultContractsHttpHandler(options DefaultContractsHttpHandlerOptions) (*DefaultContractsHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultContractsHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultContractsHttpHandler struct {
	adapter input.ContractInputAdapter
}

type DefaultContractsHttpHandlerOptions struct {
	Adapter input.ContractInputAdapter
}

func (handler *DefaultContractsHttpHandler) RegisterContractRoutes(r *mux.Router) {
	//register endpoints

	r.HandleFunc("/api/projects/{project}/contracts", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateContractRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateContract(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateContractResponse)
	}).Methods("POST")
}
