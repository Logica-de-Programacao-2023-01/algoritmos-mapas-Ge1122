package main

import "fmt"
func main() {
	expenses := map[string]float64{
		"pedro": 30.0,
		"jessica":  15.0,
		"jullia": 18.0,
	}

	equalizedExpenses := equalizeExpenses(expenses)

	fmt.Println("Equalized Expenses:")
	for person, amount := range equalizedExpenses {
		fmt.Printf("%s: %.2f\n", person, amount)
	}
}

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	total := 0
	pessoas := len(expenses)
	equalAmount := total / float64(pessoas)
	result := make(map[string]float64)

	for _, expense := range expenses {
		total += expense
	}

	for person, expense := range expenses {
		amount := expense - equalAmount
		result[person] = amount
	}

	return result
}

