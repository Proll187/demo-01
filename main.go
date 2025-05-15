package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Print("__Калькулятор индекса массы тела__ \n")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
	if IMT < 16 {
		fmt.Println("У вас сильный дефицит массы тела")
	} else if IMT < 18.5 {
		fmt.Println("У вас дефецит массы тела")
	} else if IMT < 25 {
		fmt.Println("У вас нормальный вес")
	} else if IMT < 30 {
		fmt.Println("У вас избыточный вес")
	} else {
		fmt.Println("У вас ожирение")
	}
}
func outputResult(imt float64){
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}
func calculateIMT(userKg float64, userHeight float64) float64{
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
 	return IMT
}
func getUserInput() (float64, float64){
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}