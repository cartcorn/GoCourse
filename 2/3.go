package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a = a + b
			a, b = b, a
		}()
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
