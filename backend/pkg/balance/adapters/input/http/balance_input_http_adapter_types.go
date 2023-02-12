package balanceInputHttpAdapter

type Balance struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	CurrentYear      int     `json:"currentYear"`
	PreviousYear     int     `json:"previousYear"`
	CurrentAmount    float64 `json:"currentAmount"`
	PreviousAmount   float64 `json:"previousAmount"`
	DifferenceAmount float64 `json:"differenceAmount"`
	Increase         int     `json:"increase"`
	Map              map[string]Balance
}

type Type struct {
	Balance
	Kinds []Kind `json:"kinds"`
}

type Kind struct {
	Balance
	Parents []Parent `json:"parents"`
}

type Parent struct {
	Balance
	Childs []Child `json:"childs"`
}

type Child struct {
	Balance
}

type ListBalanceRequest struct {
	Year string `json:"year"`
}

type ListBalanceResponse struct {
	Items []Type `json:"items"`
}
type ListBalanceResponseWrapper struct {
	ListBalanceResponse ListBalanceResponse
}
