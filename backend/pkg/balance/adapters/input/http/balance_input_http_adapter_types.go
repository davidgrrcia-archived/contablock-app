package balanceInputHttpAdapter

type Balance struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	CurrentYear    int                `json:"currentYear"`
	PreviousYear   int                `json:"previousYear"`
	CurrentAmount  float64            `json:"currentAmount"`
	PreviousAmount float64            `json:"previousAmount"`
	Map            map[string]Balance `json:"child"`
}

type ListBalanceRequest struct {
	Year string `json:"year"`
}

type ListBalanceResponse struct {
	Items map[string]Balance `json:"items"`
}
type ListBalanceResponseWrapper struct {
	ListBalanceResponse ListBalanceResponse
}
