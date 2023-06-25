package balanceInputHttpAdapter

type Balance struct {
	Project        string             `json:"project"`
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	CurrentYear    int                `json:"currentYear"`
	PreviousYear   int                `json:"previousYear"`
	CurrentAmount  float64            `json:"currentAmount"`
	PreviousAmount float64            `json:"previousAmount"`
	Map            map[string]Balance `json:"child"`
	Accounts       []string           `json:"accounts"`
}

type BalanceItem struct {
	Project        string        `json:"project"`
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	CurrentYear    int           `json:"currentYear"`
	PreviousYear   int           `json:"previousYear"`
	CurrentAmount  float64       `json:"currentAmount"`
	PreviousAmount float64       `json:"previousAmount"`
	Child          []BalanceItem `json:"child"`
	Accounts       []string      `json:"accounts"`
}

type ListBalanceRequest struct {
	Project string `json:"project"`
	Year    string `json:"year"`
}

type ListBalanceResponse struct {
	Items                     []BalanceItem `json:"items"`
	CurrentDebitTotalAmount   float64       `json:"currentDebitTotalAmount"`
	PreviousDebitTotalAmount  float64       `json:"previousDebitTotalAmount"`
	CurrentCreditTotalAmount  float64       `json:"currentCreditTotalAmount"`
	PreviousCreditTotalAmount float64       `json:"previousCreditTotalAmount"`
}
type ListBalanceResponseWrapper struct {
	ListBalanceResponse ListBalanceResponse
}

type CreateBalanceRequest struct {
	Project  string   `json:"project"`
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Kind     *string  `json:"kind"`
	Parent   *string  `json:"parent"`
	Child    *string  `json:"child"`
	Accounts []string `json:"accounts"`
}

type CreateBalanceResponse struct {
	Project  string   `json:"project"`
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Kind     *string  `json:"kind"`
	Parent   *string  `json:"parent"`
	Child    *string  `json:"child"`
	Accounts []string `json:"accounts"`
}

type CreateBalanceResponseWrapper struct {
	CreateBalanceResponse CreateBalanceResponse
}
