package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	//var agrem string
	for {
		height, kg := getUserInput()
		imt, err := CalcIMT(height, kg)
		if err != nil {
			fmt.Println(err)
			continue
			//panic("Неверные параметры")
		}
		outputResult(imt)
		isAgrem, err := Agrement()
		if err != nil {
			fmt.Println(err)
			continue
			//panic("Неверные данные")

		}

		//fmt.Println(err)
		//	break
		//}
		if !isAgrem {
			break

		}

	}

}

func Agrement() (bool, error) {
	fmt.Println("Желаете провести еще расчет? Yes/No")
	var agrem string
	fmt.Scan(&agrem)
	if agrem != "Yes" || agrem != "yes" {
		return false, errors.New("Введите Yes или No")

	}
	if agrem == "Yes" || agrem == "yes" {
		return true, nil
	}
	fmt.Print("Мы закончили с расчетами")
	return false, nil

}
func outputResult(imt float64) {
	result := fmt.Sprintf("%s %.2f", "Ваш ИМТ:", imt)
	switch {
	case imt < 16:
		fmt.Println("Иди съешь шавуху")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальый вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("Ты слишком жирный")
	}
	fmt.Println(result)

}
func CalcIMT(height float64, kg float64) (float64, error) {
	if kg <= 0 || height <= 0 {
		return 0, errors.New("Неверно указан вес или рост")

	}
	const power = 2
	imt := kg / math.Pow(height/100, power)
	return imt, nil

}
func getUserInput() (float64, float64) {
	var height, kg float64

	fmt.Print("Введите свой рост: ")
	fmt.Scan(&height)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&kg)
	return height, kg
}
