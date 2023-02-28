package userInputHttpAdapter

type User struct {
	Project      string `json:"project"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	EthAddress   string `json:"ethAddress"`
	CreationDate int64  `json:"creationDate"`
	LastUpdate   int64  `json:"lastUpdate"`
}

type DescribeUserRequest struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}
type DescribeUserResponse struct {
	User
}
type DescribeUserResponseWrapper struct {
	DescribeUserResponse DescribeUserResponse
}

type ListUserRequest struct {
	Project string `json:"project"`
}

type ListUserResponse struct {
	Items []User `json:"items"`
}
type ListUserResponseWrapper struct {
	ListUserResponse ListUserResponse
}

type CreateUserRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	EthAddress string `json:"ethAddress"`
}

type CreateUserResponse struct {
	User
}

type CreateUserResponseWrapper struct {
	CreateUserResponse CreateUserResponse
}
