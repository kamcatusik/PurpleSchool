package verify

import (
	"3-validation-api/configs"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailHandler struct {
	*configs.Config
}

func NewEmailHandler(router *http.ServeMux, conf EmailHandler) {
	handler := &EmailHandler{
		Config: conf.Config,
	}
	router.HandleFunc("POST /send", handler.send)
	router.HandleFunc("/verify", handler.verify)
}
func (handler *EmailHandler) send(w http.ResponseWriter, req *http.Request) {
	e := email.NewEmail()
	e.HTML = []byte("<h1>Новая отправка</h1>")
	e.Send("smtp-mail.outlook.com:587", smtp.PlainAuth("", handler.Email, handler.Password, "smtp-mail.outlook.com"))
}
func (handler *EmailHandler) verify(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("verify"))
}
