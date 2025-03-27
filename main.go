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

var calculate = map[string]func(arr []int) float64{
	"AVG": AVG,
	"SUM": SUM,
	"MED": MED,
}

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
		menuFunc := calculate[operation]
		if menuFunc == nil {
			break
		}
		fmt.Println(menuFunc(arr))
		break
	}

	/*switch operation {
	case "AVG":
		fmt.Println(AVG(arr))
	case "SUM":
		fmt.Println(SUM(arr))
	case "MED":
		fmt.Println(MED(arr))
	}
	*/

}

// Расчет среднего арифмитического
func AVG(arr []int) float64 {
	var avg, sum float64

	for _, elem := range arr {
		sum += float64(elem)
	}
	avg = sum / float64(len(arr))
	return avg
}

// Расчет суммы массива
func SUM(arr []int) float64 {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return float64(sum)
}

// расчет медианы
func MED(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	if n%2 == 0 {
		return AVG(arr)
	} else {
		return float64(arr[n/2])
	}
}

// обработка ввода
func input() (string, []int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию SUM/MED/AVG")
	scanner.Scan()
	operation := scanner.Text()
	UpOperation := strings.ToUpper(operation)
	if UpOperation != "SUM" && UpOperation != "MED" && UpOperation != "AVG" {
		return "", nil, errors.New("неверный ввод, введите SUM/MED/AVG")
	}

	fmt.Println("Введите числа через запятую")
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
