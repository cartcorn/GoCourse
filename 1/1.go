package main

import "fmt"

func main() {
	const usd float64 = 75.77
	var rub float64
	fmt.Println("Введите сумму в рублях:")
	fmt.Scanln(&rub)
	x := rub / usd
	fmt.Println("Ваша сумма в долларах США: ", x)

}
