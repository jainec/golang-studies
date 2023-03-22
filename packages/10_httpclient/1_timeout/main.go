package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	res, err := c.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
