package accountInputHttpAdapter

type Account struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Level string  `json:"level"`
	Type  *string `json:"type"`
}

type DescribeAccountRequest struct {
	ID string `json:"id"`
}
type DescribeAccountResponse struct {
	Account
}
type DescribeAccountResponseWrapper struct {
	DescribeAccountResponse DescribeAccountResponse
}

type ListAccountRequest struct{}

type ListAccountResponse struct {
	Items []DescribeAccountResponse `json:"items"`
}
type ListAccountResponseWrapper struct {
	ListAccountResponse ListAccountResponse
}
