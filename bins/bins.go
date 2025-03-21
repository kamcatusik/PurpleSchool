package bins

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
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
func CreatBinList(r *Bin) BinList {
	var bin BinList
	var name string
	var privat bool
	if r != nil {
		bin = append(bin, *r)
	}
	for {
		//спрашиваем о новом элемента
		fmt.Println("Добавить новый элемент в список yes/no")
		fmt.Scan(&name)
		if name == "no" {
			break
		}
		//добавлем имя нового элемента
		name, err := Input("Введите название")
		if err != nil {
			fmt.Println(err)
			continue
		}
		//приватизируем наш БИН
		privatStr, err := Input("Введите true или false")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if privatStr == "false" {
			privat = false
		} else {
			privat = true
		}
		//создаем Бин
		newBin1 := NewBin(name, privat)
		//добавляем в список наш бин
		bin = append(bin, *newBin1)
	}
	//возвращаем список бинов
	return bin
}
func Input(s string) (string, error) {
	var result string
	fmt.Println(s)
	_, err := fmt.Scanln(&result)
	if err != nil || result == "" {
		return "", errors.New("введены не коректные данные, повторите ввод")
	}

	return result, nil
}
