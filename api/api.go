package api

import (
	"bytes"
	"cli/jason/bins"
	"cli/jason/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiResp struct {
	Record   json.RawMessage `json:"record"`
	RespData struct {
		Id        string    `json:"id"`
		Private   bool      `json:"private"`
		CreatedAt time.Time `json:"createdAt"`
		Name      string    `json:"name"`
	} `json:"metadata"`
}
type UpdateResp struct {
	Record         json.RawMessage `json:"record"`
	UpdateRespData struct {
		ParentId string `json:"parentId"`
		Private  bool   `json:"private"`
	} `json:"metadata"`
}

const baseUrl = "https://api.jsonbin.io/v3/b/"

// запроса функция
func requestApi(method, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	//делаем запрос
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("не создан запрос")
	}
	fmt.Println("Бин верный1")
	masterKey := config.NewConfig().MasterKey
	//req.Header.Set()
	if masterKey == "" {
		return nil, fmt.Errorf("отсутствует X_MASTER_KEY")
	}

	req.Header.Set("X-Master-Key", masterKey)
	for key, value := range header {
		req.Header.Set(key, value)

	}
	//отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("запрос не отправлен")
	}

	if resp.StatusCode != 200 {

		return nil, errors.New(resp.Status)

	}
	fmt.Println("Бин верный3")
	return resp, nil
}

func CreateBinPost(data []byte, nameBin string) (*bins.Bin, error) {
	headers := map[string]string{
		"X-Bin-Private": "true",
		"X-Bin-Name":    nameBin,
		"Content-Type":  "application/json",
	}

	resp, err := requestApi("POST", baseUrl, bytes.NewBuffer(data), headers)
	if err != nil {

		return nil, errors.New("запрос не выполнен")
	}

	if resp == nil {

		return nil, errors.New("получен пустой ответ")
	}
	defer resp.Body.Close()

	//создае структуру для хранения
	var result ApiResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, errors.New("не удалось прочитать ответ")
	}

	json.Unmarshal(body, &result)

	return &bins.Bin{
		Id:        result.RespData.Id,
		Private:   result.RespData.Private,
		CreatedAt: result.RespData.CreatedAt,
		Name:      result.RespData.Name,
	}, nil
}

func UpdateBin(data []byte, id string) (*UpdateResp, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requestApi("PUT", baseUrl+id, bytes.NewBuffer(data), headers)
	if err != nil {
		fmt.Printf("ошибка: %v", err)
		return nil, fmt.Errorf("ошибка: %v", err)
	}
	fmt.Println("Бин верный")
	defer resp.Body.Close()
	var result UpdateResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, errors.New("не удалось прочитать ответ")
	}

	json.Unmarshal(body, &result)

	return &result, nil
}
func GetBin(id string) (*bins.Bin, error) {
	resp, err := requestApi("GET", baseUrl+id, nil, nil)
	if err != nil {

		return nil, errors.New("не удалось запросить Бин")
	}
	defer resp.Body.Close()
	var result ApiResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, errors.New("не удалось прочитать ответ")
	}
	json.Unmarshal(body, &result)

	return &bins.Bin{
		Id:        result.RespData.Id,
		Private:   result.RespData.Private,
		CreatedAt: result.RespData.CreatedAt,
		Name:      result.RespData.Name,
	}, nil
}

func DeleteBin(id string) error {
	resp, err := requestApi("DELETE", baseUrl+id, nil, nil)

	if err != nil {

		return errors.New("не удалось  удалить бин")
	}
	resp.Body.Close()
	return nil
}
