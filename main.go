package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func createSlice(chCreate chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		arr = append(arr, n)

	}

	for _, elem := range arr {
		chCreate <- elem
	}

}
func squer(chSquer chan int, chCreate chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		res := <-chCreate

		double := res * res

		chSquer <- double
	}

}

func main() {
	var wg sync.WaitGroup
	chCreate := make(chan int)
	chSquer := make(chan int)
	wg.Add(2)
	go createSlice(chCreate, &wg)
	go squer(chSquer, chCreate, &wg)
	go func() {
		wg.Wait()
		close(chCreate)
		close(chSquer)
	}()
	for elem := range chSquer {
		fmt.Printf("%d ", elem)
	}

}
