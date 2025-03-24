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

// создаем масив Бинов
type BinList []Bin

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
	//инициализируем новую хуету по структуре
	newBin := &Bin{

		Private:   privat,
		CreatedAt: time.Now(),
		Name:      name,
	}
	//генерим ИД
	newBin.genID()
	return newBin
}

// создаем список Бинов
func CreatBinList(bins *Bin) BinList {
	var bin BinList

	var privat bool
	file, err := files.ReadFile("save.json")
	if bin != nil {
		bin = append(bin, *bins)
	}

	newBin1 := NewBin(name, privat)
	//добавляем в список наш бин
	bin = append(bin, *newBin1)
	//возвращаем список бинов
	return bin
}
