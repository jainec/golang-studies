package main

import "fmt"

type (
	Person interface {
		Greet()
	}

	Teacher struct {
		Name string
	}

	Student struct {
		Name string
	}
)

func (t Teacher) Greet() {
	fmt.Printf("Hello, I am teacher %s\n", t.Name)
}

func (s Student) Greet() {
	fmt.Printf("Hello, I am student %s\n", s.Name)
}

func MakeGreet(people ...Person) {
	for _, person := range people {
		person.Greet()
	}
}

func main() {
	student := Student{Name: "Jaine"}
	teacher := Teacher{Name: "Adrielle"}

	MakeGreet(student, teacher)
}
