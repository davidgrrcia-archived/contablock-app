package contractdb

type ContractDb struct {
	Project      string `storage:"project"`
	ID           string `storage:"id"`
	Name         string `storage:"name"`
	Description  string `storage:"description"`
	EthAddress   string `storage:"eth_address"`
	Network      string `storage:"network"`
	AdminID      string `storage:"admin_id"`
	CreationDate string `storage:"creation_date"`
	LastUpdate   string `storage:"last_update"`
}
