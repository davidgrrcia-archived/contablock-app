package entryInputHttpAdapter

type Entry struct {
	Project         string  `json:"project"`
	ID              string  `json:"id"`
	ContractID      string  `json:"contractId"`
	UserID          string  `json:"userId"`
	AccountID       string  `json:"accountId"`
	MultipleEntryID string  `json:"multipleEntryId"`
	BalanceByYearID *string `json:"balanceByYearId"`
	Currency        string  `json:"currency"`
	Concept         *string `json:"concept"`
	Date            int64   `json:"date"`
	Quarter         int     `json:"quarter"`
	DebitBalance    float64 `json:"debitBalance"`
	CreditBalance   float64 `json:"creditBalance"`
	TotalBalance    float64 `json:"totalBalance"`
	State           string  `json:"state"`
	StateReason     *string `json:"stateReason"`
	CreationDate    int64   `json:"creationDate"`
	LastUpdate      int64   `json:"lastUpdate"`
}

type CreateEntryRequest struct {
	Project         string   `json:"project"`
	ID              *string  `json:"id"`
	ContractID      string   `json:"contractId"`
	UserID          string   `json:"userId"`
	AccountID       string   `json:"accountId"`
	MultipleEntryID string   `json:"multipleEntryId"`
	Currency        string   `json:"currency"`
	Concept         *string  `json:"concept"`
	Date            int64    `json:"date"`
	DebitBalance    *float64 `json:"debitBalance"`
	CreditBalance   *float64 `json:"creditBalance"`
}

type CreateEntryResponse struct {
	Entry
}
type CreateEntryResponseWrapper struct {
	CreateEntryResponse CreateEntryResponse
}

type DescribeEntryRequest struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}
type DescribeEntryResponse struct {
	Entry
}
type DescribeEntryResponseWrapper struct {
	DescribeEntryResponse DescribeEntryResponse
}

type ListEntryRequest struct {
	Project string `json:"project"`
}

type ListEntryResponse struct {
	Items []DescribeEntryResponse `json:"items"`
}
type ListEntryResponseWrapper struct {
	ListEntryResponse ListEntryResponse
}

type DeleteEntryRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}
type DeleteEntryResponse struct {
	Entry
}
type DeleteEntryResponseWrapper struct {
	DeleteEntryResponse DeleteEntryResponse
}

type PublishEntryRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}
type PublishEntryResponse struct {
	Entry
}
type PublishEntryResponseWrapper struct {
	PublishEntryResponse PublishEntryResponse
}

type ReviewEntryRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}
type ReviewEntryResponse struct {
	Entry
}
type ReviewEntryResponseWrapper struct {
	ReviewEntryResponse ReviewEntryResponse
}

type ApproveEntryRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}
type ApproveEntryResponse struct {
	Entry
}
type ApproveEntryResponseWrapper struct {
	ApproveEntryResponse ApproveEntryResponse
}

type RejectEntryRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}
type RejectEntryResponse struct {
	Entry
}
type RejectEntryResponseWrapper struct {
	RejectEntryResponse RejectEntryResponse
}

type UpdateEntryRequest struct {
	Project string  `json:"project"`
	ID      string  `json:"id"`
	State   *string `json:"state"`
}

type UpdateEntryResponse struct {
	Entry
}

type UpdateEntryResponseWrapper struct {
	UpdateEntryResponse UpdateEntryResponse
}
