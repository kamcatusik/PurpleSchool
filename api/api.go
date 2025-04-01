package api

import (
	"bytes"
	"cli/jason/bins"
	"cli/jason/config"
	"cli/jason/logger"
	"encoding/json"
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

const baseUrl = "https://api.jsonbin.io/v3/b/"

// запроса функция
func requestApi(method, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	//делаем запрос
	req, err := http.NewRequest(method, url, body)
	if err != nil {

		return nil, fmt.Errorf("не создан запрос")
	}
	masterKey := config.NewConfig().MasterKey

	//req.Header.Set()

	req.Header.Set("X-Master-Key", masterKey)
	for key, value := range header {
		req.Header.Set(key, value)
		logger.InfoLog.Print(key, " ", value)
	}

	//отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		return nil, fmt.Errorf("запрос не отправлен")
	}

	if resp.StatusCode != 200 {

		return nil, fmt.Errorf("ошибка доступа: %v", resp.Status)

	}

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
		return nil, fmt.Errorf("запрос не выполнен")
	}
	defer resp.Body.Close()

	//создае структуру для хранения
	var result ApiResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return nil, fmt.Errorf("не удалось прочитать ответ")
	}
	json.Unmarshal(body, &result)

	return &bins.Bin{
		Id:        result.RespData.Id,
		Private:   result.RespData.Private,
		CreatedAt: result.RespData.CreatedAt,
		Name:      result.RespData.Name,
	}, nil
}

func UpdateBin(data []byte, id string) error {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requestApi("PUT", baseUrl+id, bytes.NewBuffer(data), headers)
	if err != nil {

		return fmt.Errorf("не удалось обновить Бин")
	}
	resp.Body.Close()
	return nil
}
func GetBin(id string) error {
	resp, err := requestApi("GET", baseUrl+id, nil, nil)
	if err != nil {

		return fmt.Errorf("не удалось запросить Бин")
	}
	defer resp.Body.Close()
	var result ApiResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return fmt.Errorf("не удалось прочитать ответ")
	}
	json.Unmarshal(body, &result)

	fmt.Printf("Данные Bin: %s\n %s", id, string(result.Record))
	return nil
}

func DeleteBin(id string) error {
	resp, err := requestApi("DELETE", baseUrl+id, nil, nil)

	if err != nil {

		return fmt.Errorf("не удалось  удалить бин")
	}
	resp.Body.Close()
	return nil
}
