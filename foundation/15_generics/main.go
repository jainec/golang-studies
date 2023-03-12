package main

type MyNumber int

// constraint
type Number interface {
	~int | ~float64
}

// func SumInt()
// func SumFloat()

func Sum[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m1 := map[string]int{"Jaine": 1000, "Adrielle": 2000, "Billy": 3000}
	m2 := map[string]float64{"Jaine": 1000.12, "Adrielle": 20.40, "Billy": 33.57}
	m3 := map[string]MyNumber{"Jaine": 1000, "Adrielle": 2000, "Billy": 3000}

	println(Sum(m1))
	println(Sum(m2))
	println(Sum(m3))

	println(compare(10, 10))
}
