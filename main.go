package main

import (
	"cli/jason/bins"
	"cli/jason/files"
	"cli/jason/logger"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	logger.LogInit()
	logger.InfoLog.Println("Программа запущена")
	defer logger.Close()

	//загружаем env файл
	err := godotenv.Load()
	if err != nil {
		logger.ErrorLog.Fatal("Не удалось найти env файл")
	}
	bin, err := bins.CreatBinList(files.NewJson("save.json"))
	if err != nil {
		logger.ErrorLog.Fatalf("Не удалось создать бины %v", err)
	}
	CreatBin(bin)
	logger.InfoLog.Println("Программа завершена")

}

// создаем новый бин
func CreatBin(bin *bins.BinListWithStor) {
	privat := false
	var name string
	var err error
	//заправшивем имя Бина
	for {
		name, err = Input("Введите название вашеего Bin")
		if err != nil {
			logger.ErrorLog.Print("Ошибка неверно задано имя")
			continue

		}
		break
	}

	privatStr, err := Input("Введите приватность вашего Bin true/false")
	if err != nil {
		logger.InfoLog.Print("Приватность ключа false")

	}

	if privatStr == "true" {
		privat = true
	}
	NewBin := bins.NewBin(name, privat)
	bin.AddBinToFile(*NewBin)
	logger.InfoLog.Print("Список Бинов создан успешно")

}
func Input(s string) (string, error) {
	var result string
	fmt.Println(s)

	_, err := fmt.Scanln(&result)
	if err != nil && result == "" {
		logger.ErrorLog.Print("Ошибка ввода или пустая строка")
		//fmt.Println("пустая строка повторите ввод")
		return "", err

	}

	return result, nil
}
