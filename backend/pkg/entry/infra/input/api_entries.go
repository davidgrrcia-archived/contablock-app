package input

import (
	"context"
	"encoding/json"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultEntriesHttpHandler(options DefaultEntriesHttpHandlerOptions) (*DefaultEntriesHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultEntriesHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultEntriesHttpHandler struct {
	adapter input.EntryInputHttpAdapter
}

type DefaultEntriesHttpHandlerOptions struct {
	Adapter input.EntryInputHttpAdapter
}

func (handler *DefaultEntriesHttpHandler) RegisterEntryRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/projects/{project}/entries", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.CreateEntryRequest
		inputPayload.Project = vars["project"]

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateEntryResponse)
	}).Methods("POST")

	r.HandleFunc("/api/projects/{project}/entries", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListEntryRequest
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptListEntries(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListEntryResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/entries/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeEntryRequest
		inputPayload.ID = vars["id"]
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptDescribeEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeEntryResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/entries/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DeleteEntryRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		inputPayload.ID = vars["id"]
		inputPayload.Project = vars["project"]

		entry, err := handler.adapter.AdaptDeleteEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, entry.DeleteEntryResponse)
	}).Methods("DELETE")

	r.HandleFunc("/api/projects/{project}/entries:publish", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.PublishEntryRequest
		inputPayload.Project = vars["project"]

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		outputPayload, err := handler.adapter.AdaptPublishEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.PublishEntryResponse)
	}).Methods("POST")

	r.HandleFunc("/api/projects/{project}/entries:review", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ReviewEntryRequest
		inputPayload.Project = vars["project"]

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		outputPayload, err := handler.adapter.AdaptReviewEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ReviewEntryResponse)
	}).Methods("POST")

	r.HandleFunc("/api/projects/{project}/entries:approve", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ApproveEntryRequest
		inputPayload.Project = vars["project"]

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		outputPayload, err := handler.adapter.AdaptApproveEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ApproveEntryResponse)
	}).Methods("POST")

	r.HandleFunc("/api/projects/{project}/entries:reject", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.RejectEntryRequest
		inputPayload.Project = vars["project"]

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		outputPayload, err := handler.adapter.AdaptRejectEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.RejectEntryResponse)
	}).Methods("POST")
}
