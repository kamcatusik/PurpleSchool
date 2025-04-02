package api_test

import (
	"cli/jason/api"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateBin(t *testing.T) {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Не удалось прочитать .env")
	}
	data := []byte(`{"bin":"created"}`)
	nameBin := "testName"
	fmt.Println("Тест запущен")
	res, err := api.CreateBinPost(data, nameBin)
	fmt.Println("Бин создан")
	if err != nil {
		t.Fatalf("Пришла ошибка: %v", err)
	}
	if res == nil {
		t.Fatal("Получен nil")
	}
	if nameBin != res.Name {
		t.Errorf("Ожидалось имя: %s, получили имя: %s", nameBin, res.Name)
	}
	err = api.DeleteBin(res.Id)
	if err != nil {
		t.Errorf("Ожидалось удалить Bin: %v", err)
	}

}
func TestGetBin(t *testing.T) {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Не удалось прочитать .env")
	}
	data := []byte(`{"bin":"created"}`)
	nameBin := "testName"
	fmt.Println("Тест запущен")
	res, err := api.CreateBinPost(data, nameBin)
	fmt.Println("Бин создан")
	if err != nil {
		t.Fatalf("Пришла ошибка: %v", err)
	}
	if res == nil {
		t.Fatal("Получен nil")
	}

	result, err := api.GetBin(res.Id)
	if err != nil {
		t.Error(err)
	}
	if result.Id != res.Id {
		t.Errorf("Ожидали получить: %s,получили: %s", res.Id, result.Id)
	}
	err = api.DeleteBin(res.Id)
	if err != nil {
		t.Errorf("Ожидалось удалить Bin: %v", err)
	}
}
