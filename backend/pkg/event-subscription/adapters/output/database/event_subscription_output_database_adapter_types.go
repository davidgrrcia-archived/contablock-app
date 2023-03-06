package eventSubscriptiondb

type EventSubscriptionDb struct {
	Project      string `storage:"project"`
	ID           string `storage:"id"`
	ContractID   string `storage:"contract_id"`
	Network      string `storage:"network"`
	CreationDate string `storage:"creation_date"`
	LastUpdate   string `storage:"last_update"`
}
