package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"ID"`
	Private   bool      `json:"Private"`
	CreatedAt time.Time `json:"TimeCreated"`
	Name      string    `json:"Name"`
}

// Создаем Структуру
func NewBin(name string, privat bool, id string) *Bin {

	newBin := &Bin{
		Id:        id,
		Private:   privat,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin
}
