package balancedb

type BalanceDb struct {
	ID       string   `storage:"id"`
	Type     string   `storage:"type"`
	Kind     *string  `storage:"kind"`
	Parent   *string  `storage:"parent"`
	Child    *string  `storage:"child"`
	Accounts []string `storage:"accounts"`
}

type BalanceByYearDb struct {
	ID        string `storage:"id"`
	BalanceID string `storage:"balance_id"`
	Year      string `storage:"year"`
	Amount    string `storage:"amount"`
}
