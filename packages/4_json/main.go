package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Number: 123, Balance: 100}
	// serialize saving result
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// serialize not saving result
	json.NewEncoder(os.Stdout).Encode(account)

	raw_json := []byte(`{"n": 2, "b": 200}`)
	var Account2 Account
	err = json.Unmarshal(raw_json, &Account2)
	if err != nil {
		panic(err)
	}
	println(Account2.Balance)
}
