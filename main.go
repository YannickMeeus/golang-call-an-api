package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var bunchOfPokemon BunchOfPokemon
	json.Unmarshal(responseData, &bunchOfPokemon)

	for _, pokemon := range bunchOfPokemon.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}
}