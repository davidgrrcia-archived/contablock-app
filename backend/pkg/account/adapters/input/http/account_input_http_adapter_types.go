package accountInputHttpAdapter

type Account struct {
	Project     string  `json:"project"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Level       string  `json:"level"`
	BalanceID   *string `json:"balanceId"`
	DisplayName string  `json:"displayName"`
}

type Group struct {
	Project     string    `json:"project"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Level       string    `json:"level"`
	BalanceID   *string   `json:"balanceId"`
	DisplayName string    `json:"displayName"`
	Accounts    []Account `json:"accounts"`
}

type DescribeAccountRequest struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}
type DescribeAccountResponse struct {
	Account
}
type DescribeAccountResponseWrapper struct {
	DescribeAccountResponse DescribeAccountResponse
}

type ListAccountRequest struct {
	Project string `json:"project"`
}

type ListAccountResponse struct {
	Items []Group `json:"items"`
}
type ListAccountResponseWrapper struct {
	ListAccountResponse ListAccountResponse
}

type CreateAccountRequest struct {
	Project   string  `json:"project"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Level     string  `json:"level"`
	BalanceID *string `json:"balanceId"`
}

type CreateAccountResponse struct {
	Account
}

type CreateAccountResponseWrapper struct {
	CreateAccountResponse CreateAccountResponse
}
