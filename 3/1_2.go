package main

import "fmt"

type Car struct {
	Mark             string
	Year             int
	VolumeBag        float64
	EnginePower      bool
	WindowsOpen      bool
	VolumeBagPercent float64
}

type CarLight struct {
	Car
	PassengerCount int
}

type CarHight struct {
	Car
	Height float64
}

func main() {
	var LightCar1 = CarLight{Car: Car{"Ford", 2010, 350, false, true, 50}, PassengerCount: 5}
	var LightCar2 = CarLight{Car: Car{"Opel", 2013, 150, true, true, 20}, PassengerCount: 5}
	var HightCar1 = CarHight{Car: Car{"Zil", 2000, 1000, true, true, 120}, Height: 3.5}

	fmt.Println(LightCar1, "\n", LightCar2, "\n", HightCar1, "\n")
	//После ремонта:
	LightCar1.Year = LightCar1.Year + 2
	LightCar1.EnginePower = true
	LightCar2.Year = LightCar2.Year + 3
	HightCar1.Height = HightCar1.Height + 1
	fmt.Println(LightCar1, "\n", LightCar2, "\n", HightCar1, "\n")

}
