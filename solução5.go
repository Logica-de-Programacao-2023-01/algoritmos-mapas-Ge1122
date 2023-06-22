package main

import "fmt"
func main() {
	text := "Hello, world!"

	occurrences := contador(text)
	fmt.Print(occurrences)
}
func contador(text string) map[string]int {
	occurrences := make(map[string]int)

	for _, character := range text {
		occurrences[string(character)]++
	}

	return occurrences
}

