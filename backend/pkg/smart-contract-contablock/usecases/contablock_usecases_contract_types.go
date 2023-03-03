package contablock

type ContaBlockContractAdapter interface {
	GrantAccess(input GrantAccessAdapterInput) error
	CreateEntry(input CreateEntryAdapterInput) error
	GetEntryData(input GetEntryDataAdapterInput) (*GetEntryDataAdapterOutput, error)
	ListMultipleEntriesIds(input ListMultipleEntriesIdsAdapterInput) (*ListMultipleEntriesIdsAdapterOutput, error)
	ListAccountIds(input ListAccountIdsAdapterInput) (*ListAccountIdsAdapterOutput, error)
	PublishEntry(input PublishEntryAdapterInput) error
	RemoveEntry(input RemoveEntryAdapterInput) error
}

type GrantAccessAdapterInput struct {
	ClientCommonInput
	AdminAddress string
}

type CreateEntryAdapterInput struct {
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

type GetEntryDataAdapterInput struct {
	ClientCommonInput
	ID string
}

type GetEntryDataAdapterOutput struct {
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

type ListMultipleEntriesIdsAdapterInput struct {
	ClientCommonInput
}

type ListMultipleEntriesIdsAdapterOutput struct {
	MultipleEntriesIds []string
}

type ListAccountIdsAdapterInput struct {
	ClientCommonInput
}

type ListAccountIdsAdapterOutput struct {
	AccountIds []string
}

type PublishEntryAdapterInput struct {
	ClientCommonInput
	ID string
}

type RemoveEntryAdapterInput struct {
	ClientCommonInput
	ID string
}
