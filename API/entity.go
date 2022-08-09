package API

type Response struct {
	// Body ResponseBody `json:"rs_body"`
}

type Request struct {
	Body RequestBody `json:"rq_body"`
}
type RequestBody struct {
	Disbursement_amount int     `json:"disbursement_amount"`
	Number_of_payment   string  `json:"number_of_payment"`
	Interest_rate       float64 `json:"interest_rate"`
	Payment_frequency   int     `json:"payment_frequency"`
	Payment_unit        string  `json:"payment_unit"`
}
