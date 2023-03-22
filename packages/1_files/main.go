package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create file
	f, err := os.Create("teste.txt")
	if err != nil {
		panic(err)
	}

	// write in file
	size, err := f.Write([]byte("Writing data..."))
	// size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully with size: %d bytes\n", size)
	f.Close()

	// read from file
	file, err := os.ReadFile("teste.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// read by chunks
	file2, err := os.Open("teste.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	// remove file
	err = os.Remove("teste.txt")
	if err != nil {
		panic(err)
	}
}
