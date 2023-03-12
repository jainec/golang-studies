package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 2) * 2
	}()

	fmt.Println(total)
}

func sum(a, b int) int {
	return a + b
}
