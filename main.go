package main

import (
	"errors"
	"fmt"
<<<<<<< HEAD
)

const USDtoEUR float64 = 0.92  //стоимость доллара к евро
const USDtoRUB float64 = 89.10 //стоимость доллара к рублю
func main() {

	//EURtoRUB := USDtoRUB / USDtoEUR
	//fmt.Printf("%.2f", EURtoRUB)
	x1, x2, x3 := input() // строка число строка
	//fmt.Println(x1, x2, x3)
	fmt.Printf("%.2f", curConv(x1, x2, x3))

}

// Считываем данные пользователя
func input() (string, int, string) {
	var cur1, cur2 string
	var amount int
	for {
		fmt.Println("Введите исходную валюту USD/RUB/EUR")
		fmt.Scan(&cur1)
		err := ValidCur(cur1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Println("Введите количество валюты")
		_, err := fmt.Scan(&amount)

		if err != nil || amount <= 0 {
			fmt.Println("Некоректные данные")
			fmt.Scan(&amount)
			continue
		} else {
			break
		}

	}

	switch {
	case cur1 == "USD":
		fmt.Println("Введите валюту конвертации RUB/EUR")
	case cur1 == "EUR":
		fmt.Println("Введите валюту конвертации RUB/USD")
	case cur1 == "RUB":
		fmt.Println("Введите валюту конвертации USD/EUR")
	}
	for {
		fmt.Scan(&cur2)
		err := ValidCur(cur2)
		if err != nil || cur2 == cur1 {
			fmt.Println(err)
			continue
		}
		break
	}

	return cur1, amount, cur2

}

// читаем ошибки валюты
func ValidCur(a string) error {
	if a != "USD" && a != "EUR" && a != "RUB" {
		return errors.New("Некоретный ввод валюты, попробуйте еще раз")
	}
	return nil
}

// расчет валюты
func curConv(cur1 string, a int, cur2 string) float64 {
	//var res  float64
	switch {
	case cur1 == "USD" && cur2 == "RUB":
		return USDtoRUB * float64(a)
	case cur1 == "USD" && cur2 == "EUR":
		return USDtoEUR * float64(a)
	case cur1 == "RUB" && cur2 == "USD":
		return float64(a) / USDtoRUB
	case cur1 == "RUB" && cur2 == "EUR":
		return float64(a) / USDtoRUB * USDtoEUR
	case cur1 == "EUR" && cur2 == "RUB":
		return USDtoRUB / USDtoEUR * float64(a)
	case cur1 == "EUR" && cur2 == "USD":
		return 1 / USDtoEUR * float64(a)

	}
	return 0

}
=======
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
>>>>>>> master
