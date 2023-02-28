package input

import (
	"context"
	"encoding/json"
	"errors"

	input "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/adapters/input/http"
	utils "github.com/davidgrcia/tfg-informatica-nextjs/pkg/utils/infra/input"

	"net/http"

	"github.com/gorilla/mux"
)

func NewDefaultUsersHttpHandler(options DefaultUsersHttpHandlerOptions) (*DefaultUsersHttpHandler, error) {
	if options.Adapter == nil {
		return nil, errors.New("invalid options infra")
	}

	return &DefaultUsersHttpHandler{
		adapter: options.Adapter,
	}, nil

}

type DefaultUsersHttpHandler struct {
	adapter input.UserInputAdapter
}

type DefaultUsersHttpHandlerOptions struct {
	Adapter input.UserInputAdapter
}

func (handler *DefaultUsersHttpHandler) RegisterUserRoutes(r *mux.Router) {
	//register endpoints
	r.HandleFunc("/api/projects/{project}/users", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.ListUserRequest
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptListUsers(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.ListUserResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		vars := mux.Vars(r)
		var inputPayload input.DescribeUserRequest
		inputPayload.ID = vars["id"]
		inputPayload.Project = vars["project"]

		outputPayload, err := handler.adapter.AdaptDescribeUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		utils.Success(w, outputPayload.DescribeUserResponse)
	}).Methods("GET")

	r.HandleFunc("/api/projects/{project}/users", func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		var inputPayload input.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}
		outputPayload, err := handler.adapter.AdaptCreateUser(ctx, inputPayload)
		if err != nil {
			utils.Failure(w, err)
			return
		}

		utils.Success(w, outputPayload.CreateUserResponse)
	}).Methods("POST")
}
