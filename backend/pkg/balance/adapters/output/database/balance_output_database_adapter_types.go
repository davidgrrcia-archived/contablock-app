package balancedb

type BalanceDb struct {
	ID     string  `storage:"id"`
	Type   string  `storage:"type"`
	Kind   *string `storage:"kind"`
	Parent *string `storage:"parent"`
	Child  *string `storage:"child"`
	Year   int     `storage:"year"`
	Amount float64 `storage:"amount"`
}

type BalanceByYearDb struct {
	ID        string  `storage:"id"`
	BalanceID string  `storage:"balance_id"`
	Year      int     `storage:"year"`
	Amount    float64 `storage:"amount"`
}
