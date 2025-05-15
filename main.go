package main

import (
	"fmt"
	"math"
)
const IMTPower = 2
func main(){
	for {
			fmt.Print("__Калькулятор индекса массы тела__ \n")
			userKg, userHeight := getUserInput()
			IMT := calculateIMT(userKg, userHeight)
			outputResult(IMT)
			isRepeateCalculation := checkRepeatCalculation()
			if !isRepeateCalculation {
				break
	}
	}

}
func outputResult(imt float64){
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5: 
		fmt.Println("У вас дефецит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас ожирение")
	}
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
func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Нужно ли провести повторный расчёт? (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}