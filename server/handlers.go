package server

type CreateAccoutAndFundBody struct {
	AccountTypeId string `json:"accountTypeId" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f"`
	FundId        string `json:"fundId" example:"018ef16a-31a7-7e11-a77d-78b2eea91e2f"`
	Balance       int    `json:"balance" example:"2500000"`
}

type GetAccountAndFundResponse struct {
	AccountName string `json:"accountName" example:"Cushon ISA"`
	FundName    string `json:"fundName" example:"Cushon Equities Fund"`
	Balance     int    `json:"balance" example:"1000000"`
}
