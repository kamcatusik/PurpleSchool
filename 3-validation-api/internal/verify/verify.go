package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/files"
	"3-validation-api/pkg/reques"
	"3-validation-api/pkg/storage"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
)

type EmailHandler struct {
	*configs.Config
	*storage.EmailListWithReadWrite
}

func NewEmailHandler(router *http.ServeMux, conf EmailHandler) {
	stor, err := storage.CreateEmailListWithReadWrite(files.NewJson("save.json"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	handler := &EmailHandler{
		Config:                 conf.Config,
		EmailListWithReadWrite: stor,
	}
	router.HandleFunc("POST /send", handler.send)
	router.HandleFunc("/verify/{hash}", handler.verify)
}
func (handler *EmailHandler) send(w http.ResponseWriter, req *http.Request) {
	//читаем запрос
	mailLoad, err := reques.Decode[storage.EmailList](req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return

	}

	//валидация емейл

	err = reques.Valid(mailLoad)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(handler.EmailListWithReadWrite.Mails)
	//сохраняем в файл
	readуMail := storage.CreateEmaillist(mailLoad.Mail, storage.CreateHash())
	handler.AddEmailToFile(*readуMail)
	//отправка емейл
	verifMail := "http://localhost:8083/verify/" + readуMail.Hash
	err = reques.MailSend(mailLoad, handler.Config, verifMail)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("Успешно2")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))

}
func (handler *EmailHandler) verify(w http.ResponseWriter, req *http.Request) {
	//req.PathValue("hash")
	hash := req.PathValue("hash")
	isDel, err := handler.DelFile(hash)
	if err != nil {
		log.Println(err)
	}
	if !isDel {
		log.Println("Почта с хэшем не найдена")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isDel)
	w.Write([]byte("Verify successfuly"))
}
