package main

import (
	"fmt"
	"sort"
	"strings"
)
func main() {
	words := []string{"amor", "roma", "alegria", "regalia", "galeira", "alergia",
		"cantiga", "catinga","arara",
	}

	anagramsGroups := anagramsOccurrences(words)

	for _, anagrams := range anagramsGroups {
		fmt.Print(anagrams)
	}
}
func anagramsOccurrences(words []string) map[string][]string {
	anagramsGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramsGroups[sortedWord] = append(anagramsGroups[sortedWord], word)
	}

	return anagramsGroups
}

func sortString(term string) string {
	sorted := strings.Split(term, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

