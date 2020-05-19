package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				close(squares)
				return
			}
			squares <- x * x
		}
	}()

	for {
		square, ok := <-squares
		if !ok {
			return
		}
		fmt.Println(square)
	}
}
