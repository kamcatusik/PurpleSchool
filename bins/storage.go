package bins

import (
	"encoding/json"
	"fmt"
)

// создаем масив Бинов
type BinList struct {
	Bin []Bin `json:"Bin"`
}
type Stor interface {
	ReadFile() ([]byte, error)
	WriteFile(content []byte) error
}

type BinListWithStor struct {
	Stor Stor
	BinList
}

// создаем список Бинов
func CreatBinList(stor Stor) (*BinListWithStor, error) {

	file, err := stor.ReadFile()
	if err != nil {

		return nil, fmt.Errorf("не прочитан файл ")
	}
	var storage BinList
	err = json.Unmarshal(file, &storage)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать файл save.json")
	}
	return &BinListWithStor{
		Stor:    stor,
		BinList: storage,
	}, nil
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
func (binlist *BinListWithStor) AddBinToFile(bin Bin) error {

	binlist.Bin = append(binlist.Bin, bin)
	data, err := binlist.ToBytes()
	if err != nil {

		return err
	}

	err = binlist.Stor.WriteFile(data)
	if err != nil {

		return err
	}

	return nil
}
func (binlist *BinListWithStor) DelBin(id string) (bool, error) {
	var bins []Bin
	isDeleted := false
	for _, bin := range binlist.Bin {
		if bin.Id == id {
			isDeleted = true
		} else {

			bins = append(bins, bin)

		}

	}
	if !isDeleted {

		return false, nil
	}

	//присваевываем нашему хранилищу новые данные сохраненные в наш слайс аккаунтов и записываем в файл
	binlist.Bin = bins

	data, err := binlist.ToBytes()
	if err != nil {
		return false, fmt.Errorf("не удалось преобразовать список бинов")

	}

	err = binlist.Stor.WriteFile(data)

	if err != nil {

		return false, fmt.Errorf("не удалось записать список бинов")
	}
	return isDeleted, nil

}
