package API

import "database/sql"

func Createtable(db *sql.DB) error {

	dropBeforeCreateTable := `DROP TABLE Promotion ;`
	_, errdropBeforeCreateTable := db.Exec(dropBeforeCreateTable)
	if errdropBeforeCreateTable != nil {
		return errdropBeforeCreateTable
	}

	createTable := `CREATE TABLE Promotion (
		promotion_name  text NOT NULL COLLATE pg_catalog."C",
		description  text ,
		start_date  date ,
		end_date  date
		);`
	_, errcreateTable := db.Exec(createTable)
	if errcreateTable != nil {
		return errcreateTable
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
