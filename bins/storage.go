package bins

import (
	"cli/jason/files"
	"encoding/json"
	"fmt"
)

// создаем масив Бинов
type BinList struct {
	Bin []Bin
}

// создаем список Бинов
func CreatBinList() (*BinList, error) {

	file, err := files.ReadFile("save.json")
	if err != nil {
		return &BinList{
			Bin: []Bin{},
		}, err
	}
	var storage BinList
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println("Не удалось разобрать файл save.json")
	}
	return &storage, nil
}

// преобразование в массив байт
func (binlist *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binlist)
	if err != nil {
		return nil, err

	}
	return file, nil
}

// добавление Бина файл
func (binlist *BinList) AddBinToFile(bin Bin) {

	binlist.Bin = append(binlist.Bin, bin)
	data, err := binlist.ToBytes()
	if err != nil {

		fmt.Println("Не удалось преобразовать файл")
	}

	files.WriteFile("save.json", data)

}
