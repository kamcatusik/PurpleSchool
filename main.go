package main

import (
	"fmt"
)

const USDtoEUR float64 = 0.92  //стоимость доллара к евро
const USDtoRUB float64 = 89.10 //стоимость доллара к рублю

func main() {
	for {
		fromCur := getCurInput("Введите исходную валюту (USD, RUB, EUR): ")

		amount := getAmount("Введите количество валюты: ")

		toCur := getCurInput(fmt.Sprintf("Введите валюту конвертации (кроме %s): ", fromCur))

		res, err := curConv(fromCur, amount, toCur)
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
	var cur1 string
	for {
		fmt.Println(s)
		fmt.Scan(&cur1)

		if cur1 != "USD" && cur1 != "EUR" && cur1 != "RUB" {
			fmt.Println("Неверная валюта")
			continue
		}
		break
	}
	return cur1
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

// конвертируем в USD
func convertToUSD(currency string, amount int) (float64, error) {
	switch currency {
	case "USD":
		return float64(amount), nil
	case "EUR":
		return float64(amount) / USDtoEUR, nil
	case "RUB":
		return float64(amount) / USDtoRUB, nil
	default:
		return 0, fmt.Errorf("валюта %s не поддерживается1", currency)
	}
}

// конвертируем из USD
func convertFromUSD(currency string, usdAmount float64) (float64, error) {
	switch currency {
	case "USD":
		return usdAmount, nil
	case "EUR":
		return usdAmount * USDtoEUR, nil
	case "RUB":
		return usdAmount * USDtoRUB, nil
	default:
		return 0, fmt.Errorf("валюта %s не поддерживается2", currency)
	}
}

// конвертация через USD
func curConv(fromCur string, amount int, toCur string) (float64, error) {
	// Конвертируем в USD
	usdAmount, err := convertToUSD(fromCur, amount)
	if err != nil {
		return 0, err
	}
	// Конвертируем из USD
	result, err := convertFromUSD(toCur, usdAmount)
	if err != nil {
		return 0, err
	}

	return result, nil
}
