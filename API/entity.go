package API

type Response struct {
	Body ResponseBody `json:"rs_body"`
}

type Request struct {
	Body RequestBody `json:"rq_body"`
}
type RequestBody struct {
	Disbursement_Amount float64    `json:"disbursement_amount"`
	Number_Of_Payment   int  `json:"number_of_payment"`
	Cal_Date			string	`json:"cal_date"`
	Payment_Frequency   int     `json:"payment_frequency"`
	Payment_Unit        string  `json:"payment_unit"`
	Account_Number		int		`json:"account_number"`
}

type ResponseBody struct {
	Installment_Amount float64 `json:"installment_amount"`

}
