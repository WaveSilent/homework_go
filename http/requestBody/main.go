package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money = 1000
var bank = 0

var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err.Error())
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount
		fmt.Println("Оплата прошла успешно:", money)
	} else {
		fmt.Println("Не хватает денег на проведение оплаты")
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err.Error())
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		money -= saveAmount
		bank += saveAmount

		fmt.Println("Новое значение переменной money:", money)

		fmt.Println("Новое значение переменной bank:", bank)
	} else {
		fmt.Println("Не хватает денег, чтобы положить в копилку")
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("fail to start:", err.Error())
	}
}
