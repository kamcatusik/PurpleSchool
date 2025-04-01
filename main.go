package main

import (
	"cli/jason/api"
	"cli/jason/bins"
	"cli/jason/files"
	"cli/jason/logger"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//запускаем логгер
	logger.LogInit()
	logger.InfoLog.Println("Программа запущена")
	defer logger.Close()
	//инициализируем флаги
	create := flag.Bool("create", false, "Создание нового Бин")
	update := flag.Bool("update", false, "Обновление Бин")
	delete := flag.Bool("delete", false, "Удаление Бина")
	get := flag.Bool("get", false, "Создание нового Бин")
	list := flag.Bool("list", false, "Создание нового Бин")
	fileData := flag.String("file", "data.json", "Создание нового файла")
	id := flag.String("id", "", "Создание нового Бин")
	binName := flag.String("name", "", "Создание имени нового Бин")
	flag.Parse()

	//загружаем env файл
	err := godotenv.Load(".env")
	if err != nil {
		panic("Не удалось прочитать .env")
	}
	bin, err := bins.CreatBinList(files.NewJson("save.json"))
	if err != nil {
		fmt.Println(err)
	}

	switch {
	case *create:
		CreatBin(bin, *fileData, *binName)
	case *update:
		UpdateBin(*fileData, *id)
	case *get:
		readBin(*id)
	case *delete:
		deleteBin(bin, *id)
	case *list:
		listBins(bin)

	}

}

// создаем новый бин
func CreatBin(bin *bins.BinListWithStor, fileName, binName string) error {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("файл не был прочитан")
	}
	newBin, err := api.CreateBinPost(data, binName)
	if err != nil {
		return fmt.Errorf("бин не создан")

	}
	if binName != "" {
		newBin.Name = binName
	}

	bin.AddBinToFile(*newBin)
	return nil
}
func UpdateBin(fileName, id string) error {
	if id == "" {
		return fmt.Errorf("пустой Id.Введите id Бина")
	}
	data, err := os.ReadFile(fileName)
	if err != nil {

		return fmt.Errorf("не удалось прочитать файл")
	}
	err = api.UpdateBin(data, id)
	if err != nil {

		return fmt.Errorf("не удалось обновить Бин")
	}
	fmt.Printf("Бин обновлен успешно: %s", id)

	return nil
}
func readBin(id string) error {
	if id == "" {
		return fmt.Errorf("пустой Id.Введите id Бина")
	}
	err := api.GetBin(id)
	if err != nil {

		return fmt.Errorf("не удалось получить Бин")
	}
	return nil
}
func deleteBin(bin *bins.BinListWithStor, id string) error {
	if id == "" {
		return fmt.Errorf("пустой Id.Введите id Бина")
	}
	err := api.DeleteBin(id)
	if err != nil {

		return fmt.Errorf("не удалось получить Бин")
	}

	isDel, err := bin.DelBin(id)
	if !isDel {
		return fmt.Errorf("бинов не найдено")
	}
	if err != nil {
		return fmt.Errorf("не удалось удалить бин из файла")
	}

	return nil
}

func listBins(binlist *bins.BinListWithStor) {
	if len(binlist.Bin) == 0 {
		fmt.Println("Список пустой")
	}

	for _, bin := range binlist.Bin {
		fmt.Printf("ID: %s name: %s\n", bin.Id, bin.Name)
	}
}
