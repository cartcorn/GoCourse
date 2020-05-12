package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// open files r and w
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	t := flag.String("name", "hello", "a string")
	flag.Parse()

	w, err := os.Create(*t)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// do the actual work
	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied %v bytes\n", n)
}
