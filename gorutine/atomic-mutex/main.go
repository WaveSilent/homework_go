package main

import (
	"fmt"
	"sync"
)

//var number int = 0

// atomic - реализовано в одну слитную операцию
// var number atomic.Int64

var slice []int

// mutex - блокирует и деблокирует другие горутины
var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	fmt.Println("slice len:", len(slice))
}
