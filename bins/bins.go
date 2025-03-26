package bins

import (
	"fmt"
	"math/rand"
	"time"
)

type Bin struct {
	Id        string    `json:"ID"`
	Private   bool      `json:"Private"`
	CreatedAt time.Time `json:"TimeCreated"`
	Name      string    `json:"Name"`
}

// генерация ID
func (n *Bin) genID() {
	fmt.Println("Генерация ID")
	choiceGenPas := []rune("1234567890abcdefgetyuiooplkjhszxvnm")
	genPas := make([]rune, 15)
	for i := range genPas {
		genPas[i] = choiceGenPas[rand.Intn(len(choiceGenPas))]
	}
	n.Id = string(genPas)
}

// Создаем Структуру
func NewBin(name string, privat bool) *Bin {

	newBin := &Bin{

		Private:   privat,
		CreatedAt: time.Now(),
		Name:      name,
	}
	//генерим ИД
	newBin.genID()
	return newBin
}
