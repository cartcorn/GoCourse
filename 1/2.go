package main

import (
	"fmt"
	"math"
)

func main() {
	var cat1, cat2 float64
	fmt.Println("Введите длину катета 1 в см:")
	fmt.Scanln(&cat1)
	fmt.Println("Введите длину катета 2 в см:")
	fmt.Scanln(&cat2)
	s := cat1 * cat2 / 2
	g := math.Sqrt(cat1*cat1 + cat2*cat2)
	p := cat1 + cat2 + g
	fmt.Println("Площадь треугольника: ", s)
	fmt.Println("Гипотенуза треугольника: ", g)
	fmt.Println("Периметр треугольника: ", p)

}
