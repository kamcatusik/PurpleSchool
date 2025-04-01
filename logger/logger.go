package logger

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

// логирование ошибок
func LogInit() {
	//открываем файл, записываем в конец файла логи, если нет файла создаем его
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Не удалось открыть файл логов %v", err)
	}

	//создаем свои логгеры
	InfoLog = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Close() {
	//достаем файл и закрываем его
	if file, ok := InfoLog.Writer().(*os.File); ok {
		file.Close()
	}
}
