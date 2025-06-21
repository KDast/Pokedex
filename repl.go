package main

import (
	"strings"
)

func cleanInput(text string) []string {
	textL := strings.ToLower(text)
	newInput := strings.Fields(textL)

	return newInput
}
