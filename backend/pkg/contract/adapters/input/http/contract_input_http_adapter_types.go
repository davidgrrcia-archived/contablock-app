package contractInputHttpAdapter

type Contract struct {
	Project      string `json:"project"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	EthAddress   string `json:"ethAddress"`
	Network      string `json:"network"`
	AdminID      string `json:"adminId"`
	CreationDate int64  `json:"creationDate"`
	LastUpdate   int64  `json:"lastUpdate"`
}

type CreateContractRequest struct {
	Project     string `json:"project"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EthAddress  string `json:"ethAddress"`
	Network     string `json:"network"`
	AdminID     string `json:"adminId"`
}

type CreateContractResponse struct {
	Contract
}

type CreateContractResponseWrapper struct {
	CreateContractResponse CreateContractResponse
}
