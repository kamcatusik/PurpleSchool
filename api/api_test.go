package api_test

import (
	"bytes"
	"cli/jason/api"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func GetEnv() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Не удалось прочитать .env")
	}
}
func TestCreateBin(t *testing.T) {
	GetEnv()
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
	GetEnv()
	data := []byte(`{"bin":"created"}`)
	nameBin := "testName"
	res, err := api.CreateBinPost(data, nameBin)
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
func TestDelBin(t *testing.T) {
	GetEnv()
	data := []byte(`{"bin":"created"}`)
	nameBin := "testName"
	res, err := api.CreateBinPost(data, nameBin)
	if err != nil {
		t.Fatalf("Пришла ошибка: %v", err)
	}

	err = api.DeleteBin(res.Id)
	if err != nil {
		t.Errorf("Ожидалось удалить Bin: %v", res.Id)
	}
}
func TestUpdateBin(t *testing.T) {
	GetEnv()
	data := []byte(`{"bin":"created"}`)
	nameBin := "testName"
	res, err := api.CreateBinPost(data, nameBin)
	if err != nil {
		t.Errorf("Пришла ошибка: %v", err)
	}
	dataUpdate := []byte(`{"bin":"Update"}`)
	result, err := api.UpdateBin(dataUpdate, res.Id)
	if err != nil {
		t.Errorf("Пришла ошибка: %v", err)
	}

	body := []byte(result.Record)
	isMatch := bytes.Equal(body, dataUpdate)
	if !isMatch {
		t.Errorf("Ожидалось получить: %v,получили: %v", dataUpdate, body)
	}

	if result.UpdateRespData.ParentId != res.Id {
		t.Errorf("Ожидалось получить: %s,получили: %s", res.Id, result.UpdateRespData.ParentId)
	}
	err = api.DeleteBin(res.Id)
	if err != nil {
		t.Errorf("Ожидалось удалить Bin: %v", res.Id)
	}
}
