package accountdb

type AccountDb struct {
	ID        string `storage:"id"`
	Name      string `storage:"name"`
	Level     string `storage:"level"`
	BalanceID string `storage:"balance_id"`
}
