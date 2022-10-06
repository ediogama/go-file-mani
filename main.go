package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	length, err := f.Write([]byte("writing data in file..."))
	// length, err := f.WriteString("Hello, world!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created! Length: %d\n", length)

	f.Close()

	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}
