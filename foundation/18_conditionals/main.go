package main

func main() {
	a := 10
	b := 2

	// &&
	// ||
	if a > b {
		println("Yes")
	} else {
		println("No")
	}

	switch a {
	case 1:
		println("1")
	case 2:
		println("2")
	default:
		println("Default")
	}
}
