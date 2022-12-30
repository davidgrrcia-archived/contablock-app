package businessdb

type BusinessDb struct {
	ID                 string `storage:"id"`
	Name               string `storage:"name"`
	SocialDenomination string `storage:"social_denomination"`
	RegisteredOffice   string `storage:"registered_office"`
	Cif                string `storage:"cif"`
	Logo               string `storage:"logo"`
	EthAddress         string `storage:"eth_address"`
}
