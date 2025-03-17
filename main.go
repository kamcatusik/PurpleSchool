package main

import (
	"fmt"
	"strings"
)

type curMap = map[string]map[string]float64

func main() {

	curency := map[string]map[string]float64{
		"USD": {
			"EUR": 0.92,
			"RUB": 89.10,
		},
		"EUR": {
			"USD": 1 / 0.92,
			"RUB": 89.10 / 0.92,
		},
		"RUB": {
			"USD": 1 / 89.10,
			"EUR": 0.92 / 89.10,
		},
	}

	for {
		fromCur := getCurInput("Введите исходную валюту (USD, RUB, EUR): ")

		amount := getAmount("Введите количество валюты: ")

		toCur := getCurInput(fmt.Sprintf("Введите валюту конвертации (кроме %s): ", fromCur))

		res, err := curConv(fromCur, amount, toCur, curency)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%.2f", res)
		break
	}

}

// Считываем  валюту
func getCurInput(s string) string {
	var cur1, curUp string

	for {
		fmt.Println(s)
		fmt.Scan(&cur1)
		curUp = strings.ToUpper(cur1)
		if curUp != "USD" && curUp != "EUR" && curUp != "RUB" {
			fmt.Println("Неверная валюта")
			continue
		}
		break
	}
	return curUp
}

// Считываем количество валюты
func getAmount(s string) int {
	var amount int
	for {
		fmt.Println(s)
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("неверное количество")
			continue

		}
		break
	}
	return amount
}

// конвертация через мап
func curConv(fromCur string, amount int, toCur string, curency curMap) (float64, error) {
	insideMap, isBe := curency[fromCur] // Ищем исходную валюту в мапе, если не находим шлем на ***
	if !isBe {
		return 0, fmt.Errorf("Валюта не найдена %s", fromCur)

	}
	value, isBe := insideMap[toCur] // а тут уже смотрим по внутреней мапе нужную нам валюту и опять же если нет то посылаем иначе все норм
	if !isBe {
		return 0, fmt.Errorf("Неверная валюта для конвертации %s", toCur)
	}
	return float64(amount) * value, nil

}
