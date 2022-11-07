package businessInputHttpAdapter

type Business struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	SocialDenomination string `json:"social_denomination"`
	RegisteredOffice   string `json:"registered_office"`
	Cif                string `json:"cif"`
	Logo               string `json:"logo"`
	EthAddress         string `json:"eth_address"`
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
