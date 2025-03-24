package main

import (
	"cli/jason/bins"
	"cli/jason/file"
	"errors"
	"fmt"
)

func main() {
	var privats bool
	var name string
	data, err := file.ReadFile("save.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	bin1 := bins.NewBin(name, privats)
	newBinList := bins.CreatBinList(bin1)
	fmt.Println(bin1.Id, bin1.Private, bin1.CreatedAt.Format("15:00"), bin1.Name)
	fmt.Println(newBinList)

}
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
	NewBin := bins.CreatBinList(name, privat)
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
