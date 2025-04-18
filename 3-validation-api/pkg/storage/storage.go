package storage

import (
	"3-validation-api/logger"
	"encoding/json"
	"fmt"
)

type StorageMail struct {
	Mails []EmailList
}
type ReadWrite interface {
	ReadFile() ([]byte, error)
	WriteFile(content []byte) error
}
type EmailListWithReadWrite struct {
	StorageMail
	ReadWrite
}

// создаем список почт
func CreateEmailListWithReadWrite(stor ReadWrite) (*EmailListWithReadWrite, error) {

	file, err := stor.ReadFile()
	if err != nil {

		return nil, fmt.Errorf("не прочитан файл ")
	}
	var storage StorageMail
	err = json.Unmarshal(file, &storage)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать файл save.json")
	}
	return &EmailListWithReadWrite{
		ReadWrite:   stor,
		StorageMail: storage,
	}, nil
}

// преобразование в массив байт
func (binlist *StorageMail) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binlist)
	if err != nil {
		return nil, err

	}

	return file, nil
}

// добавление почты и хэша в список
func (emaillist *EmailListWithReadWrite) AddEmailToFile(mail EmailList) error {

	emaillist.Mails = append(emaillist.Mails, mail)
	data, err := emaillist.ToBytes()
	if err != nil {

		return err
	}
	fmt.Println(string(data))
	logger.InfoLog.Println("Получили данные для записи")
	err = emaillist.ReadWrite.WriteFile(data)
	if err != nil {
		logger.ErrorLog.Println("Ошибка записи")
		return fmt.Errorf("ошибка %v", err)
	}
	fmt.Println("Успешно7")

	return nil
}
func (emaillist *EmailListWithReadWrite) DelFile(hash string) (bool, error) {
	var mails []EmailList
	isDeleted := false
	for _, mail := range emaillist.Mails {
		if mail.Hash == hash {
			isDeleted = true
		} else {

			mails = append(mails, mail)

		}

	}
	if !isDeleted {

		return false, nil
	}

	//присваевываем нашему хранилищу новые данные сохраненные в наш слайс и записываем в файл
	emaillist.Mails = mails

	data, err := emaillist.ToBytes()
	if err != nil {
		return false, fmt.Errorf("не удалось записать файл")

	}

	err = emaillist.ReadWrite.WriteFile(data)

	if err != nil {

		return false, fmt.Errorf("не удалось записать файл")
	}
	return isDeleted, nil

}
