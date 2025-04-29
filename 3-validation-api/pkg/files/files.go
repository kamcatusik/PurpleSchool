package files

import (
	"3-validation-api/logger"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type MailStore struct {
	filename string
}

// создаем структуру нашего файла
func NewJson(name string) *MailStore {
	logger.InfoLog.Println("Зашли в структуру")
	return &MailStore{
		filename: name,
	}

}

// читаем файл и проверяем на формат json
func (stor *MailStore) ReadFile() ([]byte, error) {
	logger.InfoLog.Println("Зашли в чтение")
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
func (stor *MailStore) WriteFile(content []byte) error {
	// Проверяем/создаем директорию если нужно
	logger.InfoLog.Println("Зашли в Запись")
	dir := filepath.Dir(stor.filename)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("не удалось создать директорию: %v", err)
		}
	}
	logger.InfoLog.Println("Проверли или создали директорию")
	// Открываем файл с правильными правами
	file, err := os.OpenFile(stor.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("не удалось создать/открыть файл: %v", err)
	}
	defer file.Close()
	//записываем в файл
	_, err = file.Write(content)
	if err != nil {

		return errors.New("не удалось записать файл")
	}

	fmt.Println("Запись успешна")

	return nil
}
