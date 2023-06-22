package main

import "fmt"

func returnMap(firstMap map[int]int, secondMap map[int]int) map[int]int {
	novomap := make(map[int]int)

	for key, value := range firstMap {
		novomap[key] = value
	}
	for key, value := range secondMap {
		novomap[key] = value
	}

	return novomap
}

func main() {
	firstMap := map[int]int{
		1: 2,
		2: 2,
		3: 1,
	}
	secondMap := map[int]int{
		1: 1,
		2: 1,
		3: 1,
	}

	novomap := returnMap(firstMap, secondMap)
	fmt.Print(novomap)
}
