package balancedb

type BalanceDb struct {
	ID       string   `storage:"id"`
	Type     string   `storage:"type"`
	Kind     *string  `storage:"kind"`
	Parent   *string  `storage:"parent"`
	Child    *string  `storage:"child"`
	Accounts []string `storage:"accounts"`
	Year     string   `storage:"year"`
	Amount   string   `storage:"amount"`
}
