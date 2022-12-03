package businessInputHttpAdapter

type Business struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	SocialDenomination string `json:"socialDenomination"`
	RegisteredOffice   string `json:"registeredOffice"`
	Cif                string `json:"cif"`
	Logo               string `json:"logo"`
	EthAddress         string `json:"ethAddress"`
}

type DescribeBusinessRequest struct {
	ID string `json:"id"`
}
type DescribeBusinessResponse struct {
	Business
}
type DescribeBusinessResponseWrapper struct {
	DescribeBusinessResponse DescribeBusinessResponse
}

type ListBusinessRequest struct{}

type ListBusinessResponse struct {
	Items []DescribeBusinessResponse `json:"items"`
}
type ListBusinessResponseWrapper struct {
	ListBusinessResponse ListBusinessResponse
}
