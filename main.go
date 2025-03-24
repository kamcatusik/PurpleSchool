package main

import (
	"cli/jason/bins"
	"errors"
	"fmt"
)

func main() {
	bin := bins.CreatBinList()
	CreatBin(bin)

}

// создаем новый бин
func CreatBin(bin *bins.BinList) {
	privat := false
	//заправшивем имя Бина
	name, err := Input("Введите название вашеего Bin")
	if err != nil {
		fmt.Print(err)

	}

	privatStr, err := Input("Введите приватность вашего Bin true/false")

	if privatStr == "true" {
		privat = true
	}

	NewBin := bins.NewBin(name, privat)
	bin.AddBinToFile(*NewBin)
}
func Input(s string) (string, error) {
	var result string
	fmt.Println(s)
	_, err := fmt.Scanln(&result)
	if err != nil || result == "" {
		return "", errors.New("введены не коректные данные, повторите ввод")
	}

	return result, nil
}
