package main

import "fmt"

type ID int

var (
	j string = "valor"
	i ID     = 1
)

func main() {
	fmt.Printf("O tipo de j é: %T\n", j)
	fmt.Printf("O valor de j é: %v\n", j)

	fmt.Printf("O tipo de i é: %T\n", i)
}
