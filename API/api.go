package API

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

// type rq_body struct { การประกาศตัวแปร จะนิยมในภาษาgo เป็นแบบ camelCase*******
// 	Disbursement_amount float64 `json:"disbursement_amount"`
// 	Number_of_payment   int     `json:"number_of_payment"`
// 	Interest_rate       float64 `json:"interest_rate"`
// 	Payment_frequency   int     `json:"payment_frequency"`
// 	Payment_unit        string  `json:"payment_unit"`
// } // ประกาศตัวแปร กับเทสเคสที่ยิง postman ไม่ตรงไทป์กันง่ะ

func Api() {
	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", getdata).Methods("POST")
	r.HandleFunc("/", get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", r))

}

func getdata(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("x_request_id", "0685a3f1-ad0c-4278-8eec-9b6888642762")  ฟิค x_request_id ใน postman ไม่ต้องมาฟิคในนี้
	// w.Header().Set("x_job_id", "200330a8dd2a6b88443066") ฟิค x_request_id ใน postman ไม่ต้องมาฟิคในนี้

	// rqbody := rq_body{Payment_unit: "M"} ฟิค Payment_unit = "M" ใน postman ไม่ต้องมาฟิคในนี้
	var rqbody Request
	err := json.NewDecoder(r.Body).Decode(&rqbody)
	// _ = json.NewDecoder(r.Body).Decode(&rqbody) // return err ก็พยายามเช็คมันหน่อย มามันแปลงได้หรือไม่ได้จริงๆ
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	// rqbody := rq_body{Disbursement_amount: 35000, Number_of_payment: 4, Interest_rate: 9.12000, Payment_frequency: 1, Payment_unit: "M"}
	rqbody.Interest_rate = rqbody.Interest_rate / 100 / 12
	res := rqbody.Disbursement_amount / ((1 - (1 / (math.Pow(1+rqbody.Interest_rate, float64(rqbody.Number_of_payment))))) / rqbody.Interest_rate)
	// fmt.Fprintln(w, res)
	// json.NewEncoder(w).Encode(res)

	//Response
	response := Response{}

	json.NewEncoder(w).Encode(response)
	//  หรือ
	// js, err := json.Marshal(response)
	// if err != nil {
	// 	fmt.Fprintln(w, "error", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(js)
	// reposne  ที่ตอบกลับอยากได้ format แบบนี้
	// {
	// 	"rs_body": {
	// 		"installment_amount":
	// 	}
	// }
}
func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

/* {
    "disbursement_amount":35000,
    "number_of_payment":4,
    "interest_rate":9.12,
    "payment_frequency":1,
    "payment_unit":"M"
}
*/
