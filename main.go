package main

import (
	"fmt"
	"os"

	"github.com/siddhant-vij/HTML-Link-Parser/parser"
)

func main() {
	files := []string{
		"examples/ex1.html",
		"examples/ex2.html",
		"examples/ex3.html",
		"examples/ex4.html",
		"examples/ex5.html",
		"examples/ex6.html",
	}

	for _, file := range files {
		f, err := os.OpenFile(file, os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		links, err := parser.Parse(f)
		if err != nil {
			panic(err)
		}
		fmt.Println("File: ", file)
		fmt.Println("Links: ", links)
		fmt.Println("------------------")
	}
}
