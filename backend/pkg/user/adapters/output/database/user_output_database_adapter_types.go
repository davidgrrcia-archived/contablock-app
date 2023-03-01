package userdb

type UserDb struct {
	Project                string `storage:"project"`
	ID                     string `storage:"id"`
	Name                   string `storage:"name"`
	Email                  string `storage:"email"`
	Password               string `storage:"password"`
	Role                   string `storage:"role"`
	EthAddress             string `storage:"eth_address"`
	DefaultPasswordChanged bool   `storage:"default_password_changed"`
	Network                string `storage:"network"`
	BusinessID             string `storage:"business_id"`
	CreationDate           string `storage:"creation_date"`
	LastUpdate             string `storage:"last_update"`
}
