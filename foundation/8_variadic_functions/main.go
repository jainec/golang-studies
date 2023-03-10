package main

import "fmt"

func main() {
	fmt.Println(variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 10))
}

func variadicSum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
