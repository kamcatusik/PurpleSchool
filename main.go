package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var operation string
	var arr []int
	var err error
	for {
		operation, arr, err = input()
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	switch operation {
	case "AVG":
		fmt.Println(AVG(arr))
	case "SUM":
		fmt.Println(SUM(arr))
	case "MED":
		fmt.Println(MED(arr))
	}

}

// Расчет среднего арифмитического
func AVG(arr []int) float64 {
	var avg, sum float64

	for _, elem := range arr {
		sum += float64(elem)
	}
	avg = sum / float64((len(arr) + 1))
	return avg
}

// Расчет суммы массива
func SUM(arr []int) int {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

// расчет медианы
func MED(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	if n%2 == 0 {
		return AVG(arr)
	} else {
		return float64(n / 2)
	}
}

// обработка ввода
func input() (string, []int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию SUM/MED/AVG")
	scanner.Scan()
	operation := scanner.Text()
	UpOperation := strings.ToUpper(operation)
	if UpOperation == "SUM" || UpOperation == "MED" || UpOperation == "AVG" {

	} else {
		return "", nil, errors.New("неверный ввод, введите SUM/MED/AVG")
	}
	fmt.Println("Введител числа через запятую")
	scanner.Scan()
	input := scanner.Text()
	if len(input) == 0 {
		return "", nil, errors.New("ошибка, введите числа через запятую")
	}
	values := strings.Split(input, ",")

	var arr []int
	for i, _ := range values {
		n, err := strconv.Atoi(values[i])
		if err != nil {
			return "", nil, errors.New("ошибка при вводе, вводить числа через запятую только")
		}
		arr = append(arr, n)
	}
	return UpOperation, arr, nil
}
