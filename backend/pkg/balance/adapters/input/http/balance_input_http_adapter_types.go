package balanceInputHttpAdapter

type Balance struct {
	ID             string  `json:"id"`
	Type           string  `json:"type"`
	Kind           *string `json:"kind"`
	Parent         *string `json:"parent"`
	Child          *string `json:"child"`
	CurrentYear    *string `json:"currentYear"`
	PreviousYear   *string `json:"previousYear"`
	CurrentAmount  *string `json:"currentAmount"`
	PreviousAmount *string `json:"previousAmount"`
}

type ListBalanceRequest struct {
	Year string `json:"year"`
}

type ListBalanceResponse struct {
	Items []Balance `json:"items"`
}
type ListBalanceResponseWrapper struct {
	ListBalanceResponse ListBalanceResponse
}
