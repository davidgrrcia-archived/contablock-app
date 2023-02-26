package businessInputHttpAdapter

type Business struct {
	Project            string `json:"project"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	SocialDenomination string `json:"socialDenomination"`
	RegisteredOffice   string `json:"registeredOffice"`
	Cif                string `json:"cif"`
	Logo               string `json:"logo"`
	EthAddress         string `json:"ethAddress"`
}

type DescribeBusinessRequest struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}
type DescribeBusinessResponse struct {
	Business
}
type DescribeBusinessResponseWrapper struct {
	DescribeBusinessResponse DescribeBusinessResponse
}

type ListBusinessRequest struct {
	Project string `json:"project"`
}

type ListBusinessResponse struct {
	Items []DescribeBusinessResponse `json:"items"`
}
type ListBusinessResponseWrapper struct {
	ListBusinessResponse ListBusinessResponse
}

type CreateBusinessRequest struct {
	Project            string `json:"project"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	SocialDenomination string `json:"socialDenomination"`
	RegisteredOffice   string `json:"registeredOffice"`
	Cif                string `json:"cif"`
	Logo               string `json:"logo"`
	EthAddress         string `json:"ethAddress"`
}

type CreateBusinessResponse struct {
	Business
}

type CreateBusinessResponseWrapper struct {
	CreateBusinessResponse CreateBusinessResponse
}
