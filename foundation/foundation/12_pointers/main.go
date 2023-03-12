package main

// pointers and structs
type Account struct {
	Balance int
}

// pointers and structs
func NewAccount() *Account {
	return &Account{Balance: 0}
}

// pointers and structs
func (a *Account) simulate(value int) {
	a.Balance += value
	println(a.Balance)
}

func main() {
	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	*b = 30
	println(*b)

	// pointers and structs
	account := Account{
		Balance: 100,
	}

	account.simulate(200)
	println(account.Balance)
}
