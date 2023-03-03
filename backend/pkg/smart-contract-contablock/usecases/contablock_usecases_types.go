package contablock

const (
	// States
	EntryStateCreated = "CREATED"
)

type ClientCommonInput struct {
	ContractAddress string
	UserAddress     string
	UserPrivateKey  string
}

type GrantAccessInput struct {
	ClientCommonInput
	AdminAddress string
}

type CreateEntryInput struct {
	ClientCommonInput
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *float64
	CreditBalance   *float64
}

type GetEntryDataInput struct {
	ClientCommonInput
	ID string
}

type GetEntryDataOutput struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *float64
	CreditBalance   *float64
	State           string
}

type ListMultipleEntriesIdsInput struct {
	ClientCommonInput
}

type ListMultipleEntriesIdsOutput struct {
	MultipleEntriesIds []string
}

type ListAccountIdsInput struct {
	ClientCommonInput
}

type ListAccountIdsOutput struct {
	AccountIds []string
}

type PublishEntryInput struct {
	ClientCommonInput
	ID string
}

type RemoveEntryInput struct {
	ClientCommonInput
	ID string
}
