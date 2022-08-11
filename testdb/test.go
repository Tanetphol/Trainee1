package main


import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "postgres"
)
 type tests struct{
	disburse int
	numberpayment int
 }
func main() {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
 
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
 
    defer db.Close()
 
    // insert
    // hardcoded
    // insertStmt := `insert into "test"("Disbursement_amount", "Number_of_payment") values($1, $2)`
    // _, e := db.Exec(insertStmt, 35000, 2)
    // CheckError(e)

	// updateStmt := `update "test" set "Disbursement_amount"=$1, "Number_of_payment"=$2 where "Disbursement_amount"=$3`
	// _, e = db.Exec(updateStmt, 50000, 3, 35000)
	// CheckError(e)

	// deleteStmt := `delete from "test" where "Disbursement_amount"=$1`
	// _, e := db.Exec(deleteStmt, 50000)
	// CheckError(e)
	var r tests
	e := db.QueryRow(`SELECT "Disbursement_amount","Number_of_payment" FROM "test"`).Scan(&r.disburse,&r.numberpayment)
	CheckError(e)
	fmt.Println(r)

}
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}