package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнес газету", text, "в", i, "раз")
		time.Sleep(250 * time.Millisecond)
	}

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman(wg, "Новости")

	wg.Add(1)
	go postman(wg, "Игровой журнал")

	wg.Add(1)
	go postman(wg, "Автомобольная хроника")

	wg.Wait()

	fmt.Println("main завершился")
}
