package main

import "fmt"

func main() {
	var Fifo []int

	for i := 1; i <= 10; i++ {
		Fifo = append(Fifo, i)
	}
	fmt.Println(Fifo)
	for len(Fifo) > 0 {
		fmt.Print(Fifo[0]) // First element
		Fifo = Fifo[1:]    // Dequeue
	}

}
