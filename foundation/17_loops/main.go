package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	pets := []string{"Billy", "Felina", "Rey"}

	for k, v := range pets {
		println(k, v)
	}

	// Similar to while
	y := 0
	for y < 10 {
		println(y)
		y++
	}

	// Infinite loop
	for {
		println("Hello")
	}
}
