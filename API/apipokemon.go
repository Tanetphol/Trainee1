package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"database/sql"
)
func ApiPokemon() {
	res ,err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto")
	checkerr(err)
	var pokemon Pokemonres
	data ,err := ioutil.ReadAll(res.Body)
	checkerr(err)
	err  = json.Unmarshal(data , &pokemon)
	checkerr(err)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkerr(err)
	err = createTablePokemon(*db)
	checkerr(err)
	defer db.Close()

	for index , value := range pokemon.PokemonEntries{
		ifDataExistThenInsert(*db,index,value.EntryNumber,value.PokemonSpecies.Name,value.PokemonSpecies.Url)
	}
	
}

func checkerr(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

func createTablePokemon(db sql.DB) error{
	dropTableIfExist :=  `DROP TABLE IF EXISTS pokemon;`
	_,err := db.Exec(dropTableIfExist)
	checkerr(err)
	createTableIfNotExist := `CREATE TABLE IF NOT EXISTS pokemon(
		entry_number NUMERIC PRIMARY KEY,
		name text,
		url text);`
	_,err = db.Exec(createTableIfNotExist)
	checkerr(err)
	return nil
}
func ifDataExistThenInsert(db sql.DB,index int , entrynumber int , pokemonname string , pokemonurl string) {
	_,err := db.Query(`SELECT entry_name FROM pokemon WHERE entry_name = ?`,index+1)
	if err != nil{
		insert:= `INSERT INTO pokemon(entry_number,name,url) values($1,$2,$3)`
		_,err2 := db.Exec(insert,entrynumber,pokemonname,pokemonurl)
		checkerr(err2)
	}
	// rows, err := db.Query(`SELECT promotion_name FROM promotion WHERE '` + rqbody.Body.Cal_Date + `'  between start_date and end_date `)
}