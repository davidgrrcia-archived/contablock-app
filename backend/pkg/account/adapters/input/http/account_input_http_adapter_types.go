package accountInputHttpAdapter

type Account struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Level       string  `json:"level"`
	Type        *string `json:"type"`
	DisplayName string  `json:"displayName"`
}

type Subgroup struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Level       string    `json:"level"`
	Type        *string   `json:"type"`
	DisplayName string    `json:"displayName"`
	Accounts    []Account `json:"accounts"`
}

type Group struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Level       string     `json:"level"`
	Type        *string    `json:"type"`
	DisplayName string     `json:"displayName"`
	Subgroups   []Subgroup `json:"subgroups"`
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
	Items []Group `json:"items"`
}
type ListAccountResponseWrapper struct {
	ListAccountResponse ListAccountResponse
}
