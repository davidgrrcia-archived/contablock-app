package accountdb

import (
    "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
)

type AccountDb struct {
	ID    string `storage:"id"`
	Name  string `storage:"name"`
	Level string `storage:"level"`
	Type  string `storage:"type"`
}

func toAccount(accountDb AccountDb) (*account.Account, error) {
	var account account.Account
	account.ID = accountDb.ID
	account.Name = accountDb.Name
	account.Level = accountDb.Level
	account.Type = &accountDb.Type

	return &account, nil
}

func toAccountItems(accountSlice []AccountDb) ([]account.Account, error) {
	accountItems := make([]account.Account, len(accountSlice))
	for index := range accountSlice {
		account, err := toAccount(accountSlice[index])
		if err != nil {
			return nil, err
		}
		accountItems[index] = *account
	}
	return accountItems, nil
}
