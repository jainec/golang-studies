package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jainec/golang-studies/foundation/16_modules/math"
)

func main() {
	s := math.Sum(1, 2)
	fmt.Printf("Result: %v\n", s)
	println(math.Public)

	fmt.Println(uuid.New())
}
