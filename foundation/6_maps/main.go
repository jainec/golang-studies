package main

import "fmt"

func main() {
	salaries := map[string]int{"Jaine": 1000, "Adrielle": 2000}
	delete(salaries, "Jaine")
	salaries["Davi"] = 3000

	animals := make(map[string]int)
	colors := map[string]int{}

	animals["Billy"] = 2
	colors["Green"] = 1

	for name, salary := range salaries {
		fmt.Printf("O salario de %s é %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("O salario é %d\n", salary)
	}
}
