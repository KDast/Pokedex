package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func commandExit(cfg *config, c string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func helpMessage(cfg *config, c string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func getMap(cfg *config, c string) error {
	var body []byte // or whatever type res should be

	value, ok := cfg.cache.Get(cfg.next)
	if ok {
		body = value
		fmt.Println("getting value from cache")
	} else {
		res, err := http.Get(cfg.next)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cfg.cache.Add(cfg.next, body)
		fmt.Println("storing into cache")
	}
	var locations locationAreaEndPoint
	err := json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}
	for _, city := range locations.Results {
		fmt.Printf("%v\n", city.Name)
	}
	cfg.previous = cfg.next
	cfg.next = locations.Next

	return nil
}
func previousMap(cfg *config, c string) error {
	if cfg.previous == "" {
		fmt.Printf("no previous page available\n")
		return nil
	}

	var body []byte // or whatever type res should be

	value, ok := cfg.cache.Get(cfg.previous)
	if ok {
		body = value
		fmt.Println("getting value from cache")
	} else {
		res, err := http.Get(cfg.previous)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cfg.cache.Add(cfg.previous, body)
		fmt.Println("storing into cache")
	}

	var locations locationAreaEndPoint
	err := json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}
	for _, city := range locations.Results {
		fmt.Printf("%v\n", city.Name)
	}
	prevURL, ok := locations.Previous.(string)
	if ok {
		cfg.previous = prevURL
	} else {
		cfg.previous = ""
	}

	cfg.next = locations.Next

	return nil
}
func explore(cfg *config, city string) error {

	var body []byte
	value, ok := cfg.cache.Get(cfg.next + city)
	if ok {
		body = value
		fmt.Println("getting value from cache")
	} else {
		res, err := http.Get(cfg.next + city)
		if err != nil {
			log.Fatal(err)
			fmt.Println("invalid area")
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cfg.cache.Add(cfg.next, body)
		fmt.Println("storing into cache")
	}

	var pokemons locationAreaID
	err := json.Unmarshal(body, &pokemons)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pokemons.PokemonEncounters {
		fmt.Printf("%v\n", p.Pokemon.Name)
	}
	return nil
}
