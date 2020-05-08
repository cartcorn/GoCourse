package main

import "fmt"

func main() {
	f(c)

}

func f() {
	if x%3 == 0 {
		fmt.Println(x, "- делится без остатка")
	} else {
		fmt.Println(x, "- не делится без остатка, остаток равен: ")
	}

}
