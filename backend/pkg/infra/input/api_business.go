package input

import (
	"context"
	"errors"
	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/business"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/infra/utils"

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
	r.HandleFunc("/api/business", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		outputPayload, err := handler.adapter.AdaptListBusiness(ctx, input.ListBusinessRequest{})
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListBusinessResponse)
	}).Methods("GET")

	r.HandleFunc("/api/business/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeBusinessRequest
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptDescribeBusiness(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeBusinessResponse)
	}).Methods("GET")
}
