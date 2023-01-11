package dailyLedgerInputHttpAdapter

type DailyLedger struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Level       string  `json:"level"`
	Type        *string `json:"type"`
	DisplayName string  `json:"displayName"`
}

type ListDailyLedgerRequest struct{}

type ListDailyLedgerResponse struct {
	Items []DescribeDailyLedgerResponse `json:"items"`
}
type ListDailyLedgerResponseWrapper struct {
	ListDailyLedgerResponse ListDailyLedgerResponse
}
