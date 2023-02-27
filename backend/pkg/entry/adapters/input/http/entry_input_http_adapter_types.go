package entryInputHttpAdapter

type Entry struct {
	Project         string  `json:"project"`
	ID              string  `json:"id"`
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
	Project string `json:"project"`
	ID      string `json:"id"`
}
type DeleteEntryResponse struct {
	Entry
}
type DeleteEntryResponseWrapper struct {
	DeleteEntryResponse DeleteEntryResponse
}

type PublishEntryRequest struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}
type PublishEntryResponse struct {
	Entry
}
type PublishEntryResponseWrapper struct {
	PublishEntryResponse PublishEntryResponse
}
