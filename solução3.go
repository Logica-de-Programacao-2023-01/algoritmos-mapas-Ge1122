
package main

import "fmt"
func main() {
	values := map[int]int{
		1: 2,
		2: 2,
		3: 2,
		4: 1,
		5: 2,
	}

	sum := soma(values)
	fmt.Print(sum)
}
func soma(values map[int]int) int {
	sum := 0

	for _, valor := range values {
		sum += valor
	}

	return sum
}

