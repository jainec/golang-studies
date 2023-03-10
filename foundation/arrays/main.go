package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	println(arr[(len(arr) - 1)])

	for i, v := range arr {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
