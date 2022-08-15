package API

import "database/sql"

func Prepare(db *sql.DB) error{
	err := Createtable(db)
	if err != nil{
		return err
	}
	return nil
}
func Createtable(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE promotion ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}
	createTable := `CREATE TABLE promotion (
		promotion_name  text NOT NULL COLLATE pg_catalog."C",
		description  text ,
		start_date  date ,
		end_date  date
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}
	dropApiIfExist := `DROP TABLE IF EXISTS api;`
	_ , errdrop := db.Exec(dropApiIfExist)
	if errdrop != nil{
		return errdrop
	}
	createTableApi := `CREATE TABLE IF NOT EXISTS api(
		"Disbursement_amount" numeric,
		"Number_of_payment" integer,
		"Interest_rate" numeric,
		"Payment_frequency" integer,
		"Payment_unit" text COLLATE pg_catalog."default"
		);`
	_ , errcreate := db.Exec(createTableApi)
	if errcreate != nil{
		return errcreate
	}
	dropAccountIfExist := `DROP TABLE IF EXISTS "Account";`
	_ , errdrop = db.Exec(dropAccountIfExist)
	if errdrop != nil{
		return errdrop
	}
	createTableAccount := `CREATE TABLE IF NOT EXISTS "Account"(
		account_number NUMERIC PRIMARY KEY,
		installment_amount NUMERIC
		);`
	_ , errcreate = db.Exec(createTableAccount)
	if errcreate != nil{
		return errcreate
	}
	insert := `INSERT INTO promotion (promotion_name, description , start_date , end_date) values($1,$2,$3,$4)`
	_ , errinsert := db.Exec(insert,"Promo1","Rate < 10","2020-01-01","2020-03-31")
	if errinsert != nil {
		return errinsert
	}
	insert = `INSERT INTO promotion (promotion_name, description , start_date , end_date) values($1,$2,$3,$4)`
	_ , errinsert = db.Exec(insert,"Promo2","Rate > 10 < 20","2020-04-01","2020-06-30")
	if errinsert != nil {
		return errinsert
	}
	insert = `INSERT INTO promotion (promotion_name, description , start_date , end_date) values($1,$2,$3,$4)`
	_ , errinsert = db.Exec(insert,"Promo3","Rate > 20 < 28","2020-07-01","2020-12-30")
	if errinsert != nil {
		return errinsert
	}
	return nil

}

func CreatetableRate(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE rate ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}

	createTable := `CREATE TABLE rate (
		rate  text NOT NULL COLLATE pg_catalog."C",
		interest_rate  NUMERIC ,
		promotion_name  text );`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
	}

	return nil

}
