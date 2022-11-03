package input

import (
	"context"
	"encoding/json"
	"errors"
	input "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/game"
	"github.com/davidevloper/HoldContract-go/backend/pkg/infra/utils"

	"net/http"

	"github.com/gorilla/mux"
)

/* DEFAULT TYPES */

type DefaultGamesHttpHandler struct {
	adapter input.GameInputAdapter
}

type DefaultGamesHttpHandlerOptions struct {
	Adapter input.GameInputAdapter
}

/* CONSTRUCTOR */

func NewDefaultGamesHttpHandler(options DefaultGamesHttpHandlerOptions) (*DefaultGamesHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultGamesHttpHandler{
		adapter: options.Adapter,
	}, nil

}

/* REGISTER ROUTES FUNCTION */

func (handler *DefaultGamesHttpHandler) RegisterGameRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateGameRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		outputPayload, err := handler.adapter.AdaptCreateGame(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload.CreateGameResponse)
	}).Methods("POST")

	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		outputPayload, err := handler.adapter.AdaptListGames(ctx, input.ListGameRequest{})
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload.ListGameResponse)
	}).Methods("GET")

	r.HandleFunc("/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeGameRequest
		inputPayload.ID = vars["id"]

		game, err := handler.adapter.AdaptDescribeGame(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, game)
	}).Methods("GET")

	r.HandleFunc("/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.UpdateGameRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		vars := mux.Vars(r)
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptUpdateGame(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload.UpdateGameResponse)
	}).Methods("PUT")

	r.HandleFunc("/games/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DeleteGameRequest
		inputPayload.ID = vars["id"]

		game, err := handler.adapter.AdaptDeleteGame(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, game.DeleteGameResponse)
	}).Methods("DELETE")

	r.HandleFunc("/games/{id}:join", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.JoinGameRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		outputPayload, err := handler.adapter.AdaptJoinGame(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload.JoinGameResponse)
	}).Methods("POST")

}
