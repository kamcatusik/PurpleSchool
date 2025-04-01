package files

import (
	"cli/jason/logger"
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
		logger.ErrorLog.Printf("не удалось прчитать файл %s", stor.filename)
		return nil, fmt.Errorf("не удалось прчитать файл %s", stor.filename)
	}
	//проверка на формат json
	if !json.Valid(data) {
		logger.ErrorLog.Printf("файл %s не является валидным JSON", stor.filename)
		return nil, fmt.Errorf("файл %s не является валидным JSON", stor.filename)
	}
	logger.InfoLog.Print("Файл прочитан")
	return data, nil

}

func (stor *JsonStor) WriteFile(content []byte) error {
	//создаем файл
	file, err := os.Create(stor.filename)
	if err != nil {
		logger.ErrorLog.Print("не удалось создать файл")
		return errors.New("не удалось создать файл")
	}
	//записываем в файл
	_, err = file.Write(content)
	if err != nil {
		logger.ErrorLog.Print("не удалось записать файл")
		return errors.New("не удалось записать файл")
	}
	logger.InfoLog.Print("Файл записан")
	fmt.Println("Запись успешна")
	return nil
}
