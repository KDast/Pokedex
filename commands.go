package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func helpMessage(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func getMap(cfg *config) error {

	res, err := http.Get(cfg.next)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations locationAreaEndPoint
	err = json.Unmarshal(body, &locations)
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
func previousMap(cfg *config) error {
	if cfg.previous == "" {
		fmt.Printf("no previous page available\n")
		return nil
	}
	res, err := http.Get(cfg.previous)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locations locationAreaEndPoint
	err = json.Unmarshal(body, &locations)
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
