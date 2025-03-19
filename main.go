package main

import (
	"fmt"
	"strings"
)

type curMap = map[string]map[string]float64

var curency = curMap{
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

func main() {

	for {
		fromCur, err := getCurInput("Введите исходную валюту (USD, RUB, EUR): ") //исходная валюта

		amount := getAmount("Введите количество валюты: ")

		toCur, err := getCurInput(fmt.Sprintf("Введите валюту конвертации (кроме %s):", fromCur)) //конвертируемая валюта

		res, err := curConv(fromCur, amount, toCur, &curency) // Конвертируем
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%s: %.2f", toCur, res)
		break
	}

}

// Считываем  валюту
func getCurInput(s string) (string, error) {
	var cur1, curUp string

	for {
		fmt.Println(s)
		fmt.Scan(&cur1)
		curUp = strings.ToUpper(cur1)
		_, isBe := curency[curUp] // Ищем исходную валюту в мапе, если не находим шлем на ***
		if !isBe {
			return "", fmt.Errorf("Валюта не найдена %s", curUp)

		}
		return curUp, nil

	}

}

// Считываем количество валюты
func getAmount(s string) int {
	var amount int

	for {

		fmt.Println(s)
		_, err := fmt.Scan(&amount)

		if err != nil && amount <= 0 {
			fmt.Println("Введите целое число")
			fmt.Scanln()
			continue
		}
		break
	}
	return amount

}

// конвертация через мап
func curConv(fromCur string, amount int, toCur string, curency *curMap) (float64, error) {
	insideMap, isBe := (*curency)[fromCur] // Ищем исходную валюту в мапе, если не находим шлем на ***
	if !isBe {
		return 0, fmt.Errorf("Валюта не найдена %s", fromCur)

	}
	value, isBe := insideMap[toCur] // а тут уже смотрим по внутреней мапе нужную нам валюту и опять же если нет то посылаем иначе все норм
	if !isBe {
		return 0, fmt.Errorf("Неверная валюта для конвертации %s", toCur)
	}
	return float64(amount) * value, nil

}
