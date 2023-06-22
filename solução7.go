package main

import (
	"fmt"
	"strings"
)
func main() {
	text := "Hello, world!"

	characteristics := characteristicsSentences(text)
	fmt.Print(characteristics)
}
func characteristicsSentences(text string) map[string]map[string]int {
	words := strings.Fields(text)
	characteristics := make(map[string]map[string]int)

	for _, palavra := range words {
		letterCounter := make(map[string]int)
		for _, character := range palavra {
			letterCounter[string(character)]++
		}

		characteristics[palavra] = letterCounter
	}
	return characteristics
}

