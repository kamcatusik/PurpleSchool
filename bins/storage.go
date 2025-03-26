package bins

import (
	"encoding/json"
	"errors"
	"fmt"
)

// создаем масив Бинов
type BinList struct {
	Bin []Bin
}
type Stor interface {
	ReadFile() ([]byte, error)
	WriteFile(content []byte) error
}

type BinListWithStor struct {
	stor Stor
	BinList
}

// создаем список Бинов
func CreatBinList(stor Stor) (*BinListWithStor, error) {

	file, err := stor.ReadFile()
	if err != nil {
		return nil, err
	}
	var storage BinList
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println("Не удалось разобрать файл save.json")
	}
	return &BinListWithStor{
		stor:    stor,
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
		return errors.New("не удалось добавить новый Bin")
	}

	err = binlist.stor.WriteFile(data)
	if err != nil {
		return err
	}
	return nil
}
