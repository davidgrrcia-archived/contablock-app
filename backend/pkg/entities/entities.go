package entities

type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

type UserWallet struct {
	ID            string
	Address       string //UNIQUE
	Balance       int64
	BalanceOnHold int64
	UserID        string //UNIQUE
}

type Bet struct {
	ID     string
	Amount int64
	GameID string
	UserID string
} //UNIQUE(GameID, UserId)

type Game struct {
	ID          string
	Name        string
	OwnerUserID string
} //UNIQUE(ID, OwnerUserId)
