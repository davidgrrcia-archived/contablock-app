package input

import (
	"context"
	"encoding/json"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultBusinessHttpHandler(options DefaultBusinessHttpHandlerOptions) (*DefaultBusinessHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultBusinessHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultBusinessHttpHandler struct {
	adapter input.BusinessInputAdapter
}

type DefaultBusinessHttpHandlerOptions struct {
	Adapter input.BusinessInputAdapter
}

func (handler *DefaultBusinessHttpHandler) RegisterBusinessRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/projects/{project}/business", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListBusinessRequest
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptListBusiness(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListBusinessResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/business/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeBusinessRequest
		inputPayload.Project = vars["project"]
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptDescribeBusiness(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeBusinessResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/business", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateBusinessRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateBusiness(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateBusinessResponse)
	}).Methods("POST")
}
