package main

import "fmt"

func main() {
	var x interface{} = "Test"
	println(x.(string))

	// not panic because of 'ok'
	res, ok := x.(int)
	fmt.Printf("'res' value is %v and 'ok' value is %v\n", res, ok)

	// panic
	res2 := x.(int)
	fmt.Printf("'res2' value is: %v\n", res2)
}
