package main

import "fmt"
func main() {
	wordCounts := []map[string]int{
		{"vai,": 1, "embora!": 1},
		{"Vamos,": 1, "Jogar!": 1},
		{"jogar,": 1, "sujo!": 1},
	}

	counts := sumCounts(wordCounts)
	fmt.Print(counts)
}
func sumCounts(wordsCounts []map[string]int) map[string]int {
	counts := make(map[string]int)
	for _, count := range wordsCounts {
		for word, frequency := range count {
			counts[word] += frequency
		}
	}

	return counts
}

