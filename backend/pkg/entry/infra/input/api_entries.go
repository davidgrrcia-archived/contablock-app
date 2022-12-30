package input

import (
	"context"
	"encoding/json"
	"errors"
	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

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
	adapter input.EntryInputAdapter
}

type DefaultEntriesHttpHandlerOptions struct {
	Adapter input.EntryInputAdapter
}

func (handler *DefaultEntriesHttpHandler) RegisterEntryRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/entries", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateEntryRequest

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

	r.HandleFunc("/api/entries", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		var inputPayload input.ListEntryRequest

		outputPayload, err := handler.adapter.AdaptListEntries(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListEntryResponse)
	}).Methods("GET")

	r.HandleFunc("/api/entries/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeEntryRequest
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptDescribeEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeEntryResponse)
	}).Methods("GET")

	r.HandleFunc("/api/entries/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.UpdateEntryRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		vars := mux.Vars(r)
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptUpdateEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.UpdateEntryResponse)
	}).Methods("PUT")

	r.HandleFunc("/api/entries/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DeleteEntryRequest
		inputPayload.ID = vars["id"]

		entry, err := handler.adapter.AdaptDeleteEntry(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, entry.DeleteEntryResponse)
	}).Methods("DELETE")
}
