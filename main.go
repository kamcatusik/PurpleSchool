package main

import (
	"fmt"
	"math/rand"
)

func createSlice(chCreate chan int) {

	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)

		arr = append(arr, n)

	}

	for _, elem := range arr {
		chCreate <- elem
	}
	close(chCreate)
}
func squer(chSquer chan int, chCreate chan int) {

	for res := range chCreate {

		double := res * res

		chSquer <- double
	}
	close(chSquer)
}

func main() {

	chCreate := make(chan int)
	chSquer := make(chan int)

	go createSlice(chCreate)
	go squer(chSquer, chCreate)

	for elem := range chSquer {
		fmt.Printf("%d ", elem)
	}

}
