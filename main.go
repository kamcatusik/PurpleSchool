package main

import (
	"cli/jason/bins"
	"fmt"
)

func main() {
	var privats bool
	var name string
	//заправшивем имя Бина
	for {
		var err error
		name, err = bins.Input("Введите название вашей хуеты")
		if err != nil {
			fmt.Print(err)
			continue
		}
		break
	}

	bin1 := bins.NewBin(name, privats)
	newBinList := bins.CreatBinList(bin1)
	fmt.Println(bin1.Id, bin1.Private, bin1.CreatedAt.Format("15:00"), bin1.Name)
	fmt.Println(newBinList)

}
