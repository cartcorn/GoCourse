package main

import "fmt"

func main() {
	var c int
	fmt.Println("Введите число:")
	fmt.Scanln(&c)
	f(c)

}

func f(x int) {
	if x%2 == 0 {
		fmt.Println(x, "- это четное число")
	} else {
		fmt.Println(x, "- это нечетное число")
	}

}
