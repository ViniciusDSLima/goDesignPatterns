package main

import "fmt"

type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Nissan versa"
}

type HatchType struct {
	Name string
}

func getNewHatch() *HatchType {
	return &HatchType{}
}

func (h *HatchType) getCar() string {
	return "Fiat Palio"
}

//func main() {
//	getCarFactory(1)
//	getCarFactory(2)
//}

func getCarFactory(carType int) {

	var car Car

	if carType == 1 {
		car = getNewSedan()
	} else {
		car = getNewHatch()
	}

	carInfo := car.getCar()

	fmt.Println(carInfo)
}
