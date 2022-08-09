package API

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type rq_body2 struct {
	Disbursement_amount float64 `json:"disbursement_amount"`
	Number_of_payment   int     `json:"number_of_payment"`
	Interest_rate       float64 `json:"interest_rate"`
	Payment_frequency   int     `json:"payment_frequency"`
	Payment_unit        string  `json:"payment_unit"`
}

// var x_request_id string = "0685a3f1-ad0c-4278-8eec-9b6888642762"
// var x_job_id string ="200330a8dd2a6b88443066"
func Apiwitherr() {
	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", getdata2).Methods("POST")
	r.HandleFunc("/",get2).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getdata2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x_request_id", "0685a3f1-ad0c-4278-8eec-9b6888642762")
	w.Header().Set("x_job_id", "200330a8dd2a6b88443066")

	rqbody := rq_body2{Payment_unit: "M"}
	// rqbody := rq_body{Disbursement_amount: 35000, Number_of_payment: 4, Interest_rate: 9.12000, Payment_frequency: 1, Payment_unit: "M"}
	err := json.NewDecoder(r.Body).Decode(&rqbody)
	rqbody.Disbursement_amount = roundFloat(rqbody.Disbursement_amount,2)
	rqbody.Interest_rate = roundFloat(rqbody.Interest_rate,5)
	if err != nil {
		panic(err)
	}
	rqbody.Interest_rate = rqbody.Interest_rate / 100 / 12
	res := roundFloat(rqbody.Disbursement_amount / ((1 - (1 / (math.Pow(1+rqbody.Interest_rate, float64(rqbody.Number_of_payment))))) / rqbody.Interest_rate),2)
	json.NewEncoder(w).Encode(res)
}
func get2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"hello")
}

/* {
    "disbursement_amount":35000,
    "number_of_payment":4,
    "interest_rate":9.12,
    "payment_frequency":1,
    "payment_unit":"M"
}
*/ 
func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}
