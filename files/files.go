package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type JsonStor struct {
	filename string
}

// создаем структуру нашего файла
func NewJson(name string) *JsonStor {
	return &JsonStor{
		filename: name,
	}
}

// читаем файл и проверяем на формат json
func (stor *JsonStor) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(stor.filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прчитать файл %s", stor.filename)
	}
	//проверка на формат json
	if !json.Valid(data) {
		return nil, fmt.Errorf("файл %s не является валидным JSON", stor.filename)
	}
	return data, nil
}

func (stor *JsonStor) WriteFile(content []byte) error {
	//создаем файл
	file, err := os.Create(stor.filename)
	if err != nil {

		return errors.New("не удалось создать файл")
	}
	//записываем в файл
	_, err = file.Write(content)
	if err != nil {

		return errors.New("не удалось записать файл")
	}
	fmt.Println("Запись успешна")
	return nil
}
