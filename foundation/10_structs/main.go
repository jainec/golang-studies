package main

import "fmt"

type Address struct {
	Street string
	Number string
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) Deactivate() {
	c.Active = false
	fmt.Printf("The client %s is not active\n", c.Name)
}

func main() {
	jaine := Client{
		Name:   "Jaine",
		Age:    24,
		Active: true,
	}

	jaine.Active = false
	jaine.Address.State = "Santa Catarina"

	jaine.Deactivate()

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", jaine.Name, jaine.Age, jaine.Active)
	fmt.Printf("State: %s\n", jaine.Address.State)
}
