package main

import (
	"errors"
	"fmt"
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
		fmt.Scan(&amount)
		err := ValidCuant(amount)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
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

// ошибки ввода количества
func ValidCuant(a int) error {
	if a <= 0 {
		return errors.New("Неверные данный1")
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
