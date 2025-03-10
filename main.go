package main

import "fmt"

func main() {
	const USDtoEUR = 0.92  //стоимость дооллара к евро
	const USDtoRUB = 89.10 //стоимость доллара к рублю

	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Printf("%.2f", EURtoRUB)
}
