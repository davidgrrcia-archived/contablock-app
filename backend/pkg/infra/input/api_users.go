package input

import (
	"context"
	"encoding/json"
	"errors"
	input "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user"
	"github.com/davidevloper/HoldContract-go/backend/pkg/infra/utils"

	"net/http"

	"github.com/gorilla/mux"
)

/* DEFAULT TYPES */

type DefaultUsersHttpHandler struct {
	adapter input.UserInputAdapter
}

type DefaultUsersHttpHandlerOptions struct {
	Adapter input.UserInputAdapter
}

/* CONSTRUCTOR */

func NewDefaultUsersHttpHandler(options DefaultUsersHttpHandlerOptions) (*DefaultUsersHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultUsersHttpHandler{
		adapter: options.Adapter,
	}, nil

}

/* REGISTER ROUTES FUNCTION */

func (handler *DefaultUsersHttpHandler) RegisterUserRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		outputPayload, err := handler.adapter.AdaptCreateUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		utils.Success(w, outputPayload.CreateUserResponse)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		outputPayload, err := handler.adapter.AdaptListUsers(ctx, input.ListUserRequest{})
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload.ListUserResponse)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeUserRequest
		inputPayload.ID = vars["id"]

		user, err := handler.adapter.AdaptDescribeUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, user)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.UpdateUserRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}

		vars := mux.Vars(r)
		inputPayload.ID = vars["id"]

		outputPayload, err := handler.adapter.AdaptUpdateUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, outputPayload.UpdateUserResponse)
	}).Methods("PUT")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DeleteUserRequest
		inputPayload.ID = vars["id"]

		user, err := handler.adapter.AdaptDeleteUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
		}
		utils.Success(w, user.DeleteUserResponse)
	}).Methods("DELETE")
}
