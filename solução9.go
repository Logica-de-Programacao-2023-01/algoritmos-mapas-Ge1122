package main

import "fmt"
func main() {
	n := 10
	fibSeq := fibonacci(n)

	fmt.Println("Fibonacci Sequence:")
	for i, num := range fibSeq {
		fmt.Printf("Index %d: %d\n", i, num)
	}
}
func fibonacci(n int) map[int]int {
	sequence := make(map[int]int)

	sequence[0] = 0
	if n > 0 {
		sequence[1] = 1
	}

	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}

