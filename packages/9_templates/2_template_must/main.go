package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 80}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}}\nWorkload: {{.Workload}}\n"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
