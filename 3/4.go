package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Println(string(b))

	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("\1111.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
