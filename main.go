package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, arr := input()
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
func input() (string, []int) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите операцию SUM/MED/AVG")
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println("Введител числа через запятую")
	scanner.Scan()
	input := scanner.Text()
	values := strings.Split(input, ",")
	arr := make([]int, len(values))
	for i := range values {
		arr[i], _ = strconv.Atoi(values[i])
	}
	return operation, arr
}
