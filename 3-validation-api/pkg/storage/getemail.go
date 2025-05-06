package storage

import (
	"crypto/rand"
	"encoding/hex"
)

// возможно не хватает  кому от кого и перенести в сендемейл
type EmailList struct {
	Mail string `json:"email" validate:"required,email"`
	Hash string `json:"hash"`
}

func CreateEmaillist(mail string, hash string) *EmailList {
	NewEmail := &EmailList{
		Mail: mail,
		Hash: hash,
	}
	return NewEmail
}

func CreateHash() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
