package input

import (
	"context"
	"errors"
	input "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/bet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/infra/utils"

	"net/http"

	"github.com/gorilla/mux"
)

/* DEFAULT TYPES */

type DefaultBetHttpHandler struct {
	adapter input.BetInputAdapter
}

type DefaultBetHttpHandlerOptions struct {
	Adapter input.BetInputAdapter
}

/* CONSTRUCTOR */

func NewDefaultBetHttpHandler(options DefaultBetHttpHandlerOptions) (*DefaultBetHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultBetHttpHandler{
		adapter: options.Adapter,
	}, nil

}

/* REGISTER ROUTES FUNCTION */

func (handler *DefaultBetHttpHandler) RegisterBetRoutes(r *mux.Router) {
	//register endpoints
	//todo query string params
	r.HandleFunc("/bets/{user-id}/{game-id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		var inputPayload input.DescribeBetByUserIDRequest

		vars := mux.Vars(r)
		inputPayload.UserID = vars["user-id"]
		inputPayload.GameID = vars["game-id"]

		outputPayload, err := handler.adapter.AdaptDescribeBetByUserID(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload)
	}).Methods("GET")

	r.HandleFunc("/bets/{id}", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		var inputPayload input.DescribeBetRequest

		vars := mux.Vars(r)
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptDescribeBet(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload)

	}).Methods("GET")
}
