package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// читаем файл и проверяем на формат json
func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("не удалось прчитать файл %s", name)
	}
	if !json.Valid(data) {
		return nil, fmt.Errorf("файл %s не является валидным JSON", name)
	}
	return data, nil
}

func WriteFile(fileName string, content []byte) error {
	//создаем файл
	file, err := os.Create(fileName)
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
