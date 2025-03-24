package file

import (
	"encoding/json"
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

func WriteFile(fileName string, content []byte) {
	//создаем файл
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Не удалось создать файл")
		return
	}
	//записываем в файл
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Не удалось записать файл")
		return
	}
	fmt.Println("Запись успешна")
}
