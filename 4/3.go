package main

import (
	"fmt"
	"strconv"
)

var firstNum float32
var secondNum float32
var action string
var help string = "В калькуляторе предусмотрены функции + - * /"

func calcucate(firstNum float32, action string, secondNum float32) float32 {
	switch action {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	case "help":
		return help
	default:
		fmt.Println("Не удалось распознать действие, введите HELP для справки")
		return 0
	}

}
func inputData(msg string) (data string) {
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		panic(nil)
	}
}

func enterNumber(msg string) float32 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return float32(num)
			break
		}
		fmt.Println("Не удалось распознать число")
	}
	return 0
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	// Ввод первого числа
	firstNum := enterNumber("Введите первое число: ")
	action := inputData("Укажите действие: ")
	// Ввод второго числа
	secondNum := enterNumber("Введите второе число: ")
	fmt.Printf("Результат: %v \n", calculate(firstNum, action, secondNum))
}
