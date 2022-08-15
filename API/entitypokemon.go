package API


type Pokemonres struct{
	PokemonEntries	[]EntriesDetail `json:"pokemon_entries"`
}

type EntriesDetail struct {
	EntryNumber		int				`json:"entry_number"`
	PokemonSpecies	Pokemondetail	`json:"pokemon_species"`
}

type Pokemondetail struct {
	Name	string			`json:"name"`
	Url		string			`json:"url"`
}
// "entry_number": 1,
//             "pokemon_species": {
//                 "name": "bulbasaur",
//                 "url": "https://pokeapi.co/api/v2/pokemon-species/1/"