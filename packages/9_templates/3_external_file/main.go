package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 20},
		{"PHP", 80},
	})
	if err != nil {
		panic(err)
	}
}
