package API

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func Api() {
	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", getdata).Methods("POST")
	r.HandleFunc("/", get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", r))
}
func getdata(w http.ResponseWriter, r *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, er := sql.Open("postgres", psqlconn)
	fmt.Fprintln(w, "1")
	CheckError(er, w)
	err2 := Prepare(db)
	CheckError(err2,w)
	defer db.Close()
	var rqbody Request
	var promo string

	


	var interest_rate float64
	// var cal_date time.Time
	err := json.NewDecoder(r.Body).Decode(&rqbody)
	// cal_date = datetimetype(rqbody.Body.Cal_date)
	if err != nil {
		fmt.Fprintln(w, "1.1", err.Error())
		return
	}
	// err = db.QueryRow(`SELECT promotion_name FROM "Promotion" where ? BETWEEN start_date AND end_date`, rqbody.Body.Cal_Date).Scan(&promo)
	// fmt.Fprintln(w, "2")
	// CheckError(err, w)
	// =================================
	rows, err := db.Query(`SELECT promotion_name FROM promotion WHERE '` + rqbody.Body.Cal_Date + `'  between start_date and end_date `)
	fmt.Fprintln(w, "2")
	CheckError(err, w)
	for rows.Next() {
		if err := rows.Scan(&promo); err != nil {
			fmt.Fprintln(w,"err")
			log.Fatal(err)
		}
	}
	// =================================
	// err = db.QueryRow(`SELECT interest_rate FROM "Rate" where promotion_name = ?`, promo).Scan(&interest_rate)
	// fmt.Fprintln(w, "3")
	// fmt.Fprintln(w, interest_rate)
	// CheckError(err, w)
	// =================================
	rows2, err2 := db.Query(`SELECT interest_rate FROM "Rate" where promotion_name = '` + promo + `'`)
	CheckError(err2, w)
	for rows2.Next() {
		if err := rows2.Scan(&interest_rate); err != nil {
			fmt.Fprintln(w,"err")
			log.Fatal(err)
		}
	}
	// =================================

	interest_rate = interest_rate / 100 / 12
	res := rqbody.Body.Disbursement_Amount / ((1 - (1 / (math.Pow(1+interest_rate, float64(rqbody.Body.Number_Of_Payment))))) / interest_rate)

	insertStmt := `insert into "api"("Disbursement_amount", "Number_of_payment","Interest_rate","Payment_frequency","Payment_unit") values($1, $2,$3,$4,$5)`
	_, e := db.Exec(insertStmt, 35000, 4, 2.5, 1, "M")
	fmt.Fprintln(w, "4")
	CheckError(e, w)
	insertres := `insert into "Account"("installment_amount" , account_number) values($1,$2)`
	_, e = db.Exec(insertres, res, rqbody.Body.Account_Number)
	fmt.Fprintln(w, "5")
	CheckError(e, w)
	//Response
	response := Response{
		Body: ResponseBody{
			Installment_Amount: res,
		},
	}
	js, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintln(w, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
func CheckError(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Fprintln(w, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


// type rq_body struct { การประกาศตัวแปร จะนิยมในภาษาgo เป็นแบบ camelCase*******
// 	Disbursement_amount float64 `json:"disbursement_amount"`
// 	Number_of_payment   int     `json:"number_of_payment"`
// 	Interest_rate       float64 `json:"interest_rate"`
// 	Payment_frequency   int     `json:"payment_frequency"`
// 	Payment_unit        string  `json:"payment_unit"`
// } // ประกาศตัวแปร กับเทสเคสที่ยิง postman ไม่ตรงไทป์กันง่ะ
// w.Header().Set("x_request_id", "0685a3f1-ad0c-4278-8eec-9b6888642762")  ฟิค x_request_id ใน postman ไม่ต้องมาฟิคในนี้
// w.Header().Set("x_job_id", "200330a8dd2a6b88443066") ฟิค x_request_id ใน postman ไม่ต้องมาฟิคในนี้
// rqbody := rq_body{Payment_unit: "M"} ฟิค Payment_unit = "M" ใน postman ไม่ต้องมาฟิคในนี้
// _ = json.NewDecoder(r.Body).Decode(&rqbody) // return err ก็พยายามเช็คมันหน่อย มามันแปลงได้หรือไม่ได้จริงๆ
// fmt.Fprintln(w, res)
// json.NewEncoder(w).Encode(res)
// json.NewEncoder(w).Encode(response)
//  หรือ
// reposne  ที่ตอบกลับอยากได้ format แบบนี้
// {
// 	"rs_body": {
// 		"installment_amount":
// 	}
// }

/* {
    "disbursement_amount":35000,
    "number_of_payment":4,
    "interest_rate":9.12,
    "payment_frequency":1,
    "payment_unit":"M"
}
*/
// 	SELECT column_name(s)
// FROM table_name
// WHERE column_name BETWEEN value1 AND value2;
