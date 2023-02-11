package balanceInputHttpAdapter

type Balance struct {
	ID               string  `json:"id"`
	Type             string  `json:"type"`
	Kind             *string `json:"kind"`
	Parent           *string `json:"parent"`
	Child            *string `json:"child"`
	CurrentYear      int     `json:"currentYear"`
	PreviousYear     int     `json:"previousYear"`
	CurrentAmount    float64 `json:"currentAmount"`
	PreviousAmount   float64 `json:"previousAmount"`
	DifferenceAmount float64 `json:"differenceAmount"`
	Increase         int     `json:"increase"`
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
