package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowertext := strings.ToLower(strings.Trim(text, " "))
	formattedTexts := strings.Fields(lowertext)

	return formattedTexts
}

func main() {
	fmt.Println("Hello World!")
}
