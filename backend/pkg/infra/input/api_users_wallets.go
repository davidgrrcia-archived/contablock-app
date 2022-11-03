package input

import (
	"context"
	"encoding/json"
	"errors"
	input "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user-wallet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/infra/utils"

	"net/http"

	"github.com/gorilla/mux"
)

/* DEFAULT TYPES */

type DefaultUserWalletsHttpHandler struct {
	adapter input.UserWalletInputAdapter
}

type DefaultUserWalletsHttpHandlerOptions struct {
	Adapter input.UserWalletInputAdapter
}

/* CONSTRUCTOR */

func NewDefaultUserWalletsHttpHandler(options DefaultUserWalletsHttpHandlerOptions) (*DefaultUserWalletsHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultUserWalletsHttpHandler{
		adapter: options.Adapter,
	}, nil

}

/* REGISTER ROUTES FUNCTION */

func (handler *DefaultUserWalletsHttpHandler) RegisterUserWalletRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/users/{user-id}/wallets", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		var inputPayload input.CreateUserWalletRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		vars := mux.Vars(r)
		inputPayload.UserID = vars["user-id"]

		outputPayload, err := handler.adapter.AdaptCreateUserWallet(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload)

	}).Methods("POST")

	r.HandleFunc("/users/{user-id}/wallets", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		var inputPayload input.DescribeUserWalletsByUserIDRequest

		vars := mux.Vars(r)
		inputPayload.UserID = vars["user-id"]

		outputPayload, err := handler.adapter.AdaptDescribeUserWalletsByUserID(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload)
	}).Methods("GET")

	r.HandleFunc("/users/{user-id}/wallets/{wallet-id}", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		var inputPayload input.DescribeUserWalletRequest

		vars := mux.Vars(r)
		inputPayload.ID = vars["wallet-id"]
		inputPayload.UserID = vars["user-id"]

		outputPayload, err := handler.adapter.AdaptDescribeUserWallet(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload)

	}).Methods("GET")
}
