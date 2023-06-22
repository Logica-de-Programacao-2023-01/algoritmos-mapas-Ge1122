package main

import (
	"fmt"
	"strings"
)
func main() {
	text := "Ola, tudo bem?"

	aparições := Contador(text)
	fmt.Print(aparições)
}
func Contador(text string) map[string]int {
	words := strings.Fields(text)
	aparições := make(map[string]int)

	for _, word := range words {
		aparições[word]++
	}

	return aparições
}

