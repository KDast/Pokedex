package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v\n", input[0])

	}
	/*if err != nil {
	fmt.Fprintln(os.Stderr, "shoudln't see an error scanning a string")
	}*/
}
