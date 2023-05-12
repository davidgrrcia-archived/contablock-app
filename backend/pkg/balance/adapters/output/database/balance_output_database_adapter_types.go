package balancedb

type BalanceDb struct {
	Project  string  `storage:"project"`
	ID       string  `storage:"id"`
	Type     string  `storage:"type"`
	Kind     *string `storage:"kind"`
	Parent   *string `storage:"parent"`
	Child    *string `storage:"child"`
	Accounts string  `storage:"accounts"`
}

type BalanceByYearDb struct {
	Project   string  `storage:"project"`
	ID        string  `storage:"id"`
	BalanceID string  `storage:"balance_id"`
	Year      int     `storage:"year"`
	Amount    float64 `storage:"amount"`
}

type AddBalanceDb struct {
	Project  string  `storage:"project"`
	ID       string  `storage:"id"`
	Type     string  `storage:"type"`
	Kind     *string `storage:"kind"`
	Parent   *string `storage:"parent"`
	Child    *string `storage:"child"`
	Accounts string  `storage:"accounts"`
}
