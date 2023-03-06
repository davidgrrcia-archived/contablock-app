package eventSubscriptionInputHttpAdapter

type EventSubscription struct {
	Project      string `json:"project"`
	ID           string `json:"id"`
	ContractID   string `json:"contractId"`
	Network      string `json:"network"`
	CreationDate int64  `json:"creationDate"`
	LastUpdate   int64  `json:"lastUpdate"`
}

type CreateEventSubscriptionRequest struct {
	Project    string `json:"project"`
	ID         string `json:"id"`
	ContractID string `json:"contractId"`
	Network    string `json:"network"`
}

type CreateEventSubscriptionResponse struct {
	EventSubscription
}

type CreateEventSubscriptionResponseWrapper struct {
	CreateEventSubscriptionResponse CreateEventSubscriptionResponse
}
