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
	/* for k, v := range r.Header {
		fmt.Println("k:", k, "--v:", v)
	}

	fmt.Println("Привет", r.Header.Get("Myheader")) */

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 Внутренная ошибка сервера

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400 Некорректный http запрос

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		money -= paymentAmount
		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
	} else {
		msg := "Не хватает денег на проведение оплаты"
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	if money >= saveAmount {
		money -= saveAmount
		bank += saveAmount

		msg1 := "Новое значение переменной money:" + strconv.Itoa(money)
		fmt.Println(msg1)
		_, err := w.Write([]byte(msg1))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
			return
		}

		msg1 = "Новое значение переменной bank:" + strconv.Itoa(bank)
		fmt.Println(msg1)
		_, err = w.Write([]byte(msg1))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}

	} else {
		msg := "Не хватает денег, чтобы положить в копилку"
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err.Error())
		}
	}
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
