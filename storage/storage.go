package storage

// создаем масив Бинов
type BinList []Bin

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
