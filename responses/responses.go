package responses

type Responses struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsesFindAll struct {
	Result    interface{} `json:"result"`
	Page      string      `json:"page"`
	Limit     string      `json:"limit"`
	TotalData string      `json:"total_data"`
}

type ResponsesLogin struct {
	DataUser interface{} `json:"data_user"`
	Token    string      `json:"token"`
}

type ResultByMerchants struct {
	MerchantName     string  `json:"merchant_name"`
	Omzet            float64 `json:"omzet"`
	DateTransactions string  `json:"date_transactions"`
}

type ResultByOutlets struct {
	MerchantName     string  `json:"merchant_name"`
	OutletName       string  `json:"outlet_name"`
	Omzet            float64 `json:"omzet"`
	DateTransactions string  `json:"date_transactions"`
}
