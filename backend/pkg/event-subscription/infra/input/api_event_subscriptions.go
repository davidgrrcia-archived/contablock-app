package input

import (
	"context"
	"encoding/json"
	"errors"
	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"
	"github.com/gorilla/mux"
	"net/http"
)

func NewDefaultEventSubscriptionsHttpHandler(options DefaultEventSubscriptionsHttpHandlerOptions) (*DefaultEventSubscriptionsHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultEventSubscriptionsHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultEventSubscriptionsHttpHandler struct {
	adapter input.EventSubscriptionInputAdapter
}

type DefaultEventSubscriptionsHttpHandlerOptions struct {
	Adapter input.EventSubscriptionInputAdapter
}

func (handler *DefaultEventSubscriptionsHttpHandler) RegisterEventSubscriptionRoutes(r *mux.Router) {
	r.HandleFunc("/api/projects/{project}/event-subscriptions", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateEventSubscriptionRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateEventSubscription(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateEventSubscriptionResponse)
	}).Methods("POST")
}
