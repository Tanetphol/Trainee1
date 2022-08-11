package API

type Response struct {
	Body ResponseBody `json:"rs_body"`
}

type Request struct {
	Body RequestBody `json:"rq_body"`
}
type RequestBody struct {
	Disbursement_amount float64    `json:"disbursement_amount"`
	Number_of_payment   int  `json:"number_of_payment"`
	Cal_date			string	`json:"cal_date"`
	Payment_frequency   int     `json:"payment_frequency"`
	Payment_unit        string  `json:"payment_unit"`
	Account_number		int		`json:"account_number"`
}

type ResponseBody struct {
	Installment_amount float64 `json:"installment_amount"`

}
