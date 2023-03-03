package businessdb

type BusinessDb struct {
	Project            string `storage:"project"`
	ID                 string `storage:"id"`
	Name               string `storage:"name"`
	SocialDenomination string `storage:"social_denomination"`
	RegisteredOffice   string `storage:"registered_office"`
	Cif                string `storage:"cif"`
	Logo               string `storage:"logo"`
}
