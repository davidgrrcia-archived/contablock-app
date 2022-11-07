package businessdb

import (
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/business"
)

type BusinessDb struct {
	ID                 string `storage:"id"`
	Name               string `storage:"name"`
	SocialDenomination string `storage:"social_denomination"`
	RegisteredOffice   string `storage:"registered_office"`
	Cif                string `storage:"cif"`
	Logo               string `storage:"logo"`
	EthAddress         string `storage:"eth_address"`
}

func toBusiness(businessDb BusinessDb) (*business.Business, error) {
	var business business.Business
	business.ID = businessDb.ID
	business.Name = businessDb.Name
	business.SocialDenomination = businessDb.SocialDenomination
	business.RegisteredOffice = businessDb.RegisteredOffice
	business.Cif = businessDb.Cif
	business.Logo = businessDb.Logo
	business.EthAddress = businessDb.EthAddress

	return &business, nil
}

func toBusinessItems(businessSlice []BusinessDb) ([]business.Business, error) {
	businessItems := make([]business.Business, len(businessSlice))
	for index := range businessSlice {
		business, err := toBusiness(businessSlice[index])
		if err != nil {
			return nil, err
		}
		businessItems[index] = *business
	}
	return businessItems, nil
}
