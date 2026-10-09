package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type MessageInfo struct {
	Title     string `json:"title"`
	PostIndex int    `json:"postIndex"`
	Text      string `json:"text"`
	// Urgent - срочное
	Urgent bool `json:"urgent"`
}

var mtx sync.Mutex
var infoMessage map[int]MessageInfo = make(map[int]MessageInfo)

type HttpResponse struct {
	InfoMessage map[int]MessageInfo `json:"allInfoMessage"`
}

func saveMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "the wrong method was passed"
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	var messageInfo MessageInfo
	if err := json.NewDecoder(r.Body).Decode(&messageInfo); err != nil {
		fmt.Println("err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	randInt := rand.Intn(100)

	mtx.Lock()
	infoMessage[randInt] = messageInfo
	mtx.Unlock()

	fmt.Println("Прочитано и добавлено в список сообщение:", messageInfo, "с уникальным id:", randInt)

	mtx.Lock()
	fmt.Println("Кол-во сообщений в списке:", len(infoMessage))
	mtx.Unlock()

}

func delMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "the wrong method was passed"
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("fail to read HTTP body:", err.Error())
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		fmt.Println("fail to convert HTTP body to integer:", err.Error())
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	val, ok := infoMessage[id]
	if !ok {
		fmt.Println("Такого сообщения нет в списке!")
		return
	}
	delete(infoMessage, id)
	fmt.Println("Из списка успешно удалено сообщение:", val)

}

func AllMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "the wrong method was passed"
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	postIndexParam := r.URL.Query().Get("postIndex")

	urgentParam := r.URL.Query().Get("urgent")

	var nPointer *int

	var bPointer *bool

	if postIndexParam != "" {
		num, err := strconv.Atoi(postIndexParam)
		if err != nil {
			fmt.Println("fail to convert to integer:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		nPointer = &num
	}

	if urgentParam != "" {
		boolParam, err := strconv.ParseBool(urgentParam)
		if err != nil {
			fmt.Println("fail to convert to bool:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bPointer = &boolParam
	}

	filtred := make(map[int]MessageInfo, len(infoMessage))
	for k, v := range infoMessage {
		if nPointer != nil && *nPointer != v.PostIndex {
			continue
		}

		if bPointer != nil && *bPointer != v.Urgent {
			continue
		}

		filtred[k] = v

	}

	httpResponse := HttpResponse{
		InfoMessage: filtred,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write HTTP body", err.Error())
		return
	}

}

func findMessageHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	val, ok := infoMessage[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)

		msg := "not found message with id: " + httpRequestBodyString
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	msg := "Найдено сообщение: " + val.Text + " с id: " + httpRequestBodyString
	fmt.Println(msg)

	_, err = w.Write([]byte(msg))
	if err != nil {
		fmt.Println("fail to write HTTP body:", err.Error())
	}

}

func main() {
	http.HandleFunc("/save", saveMessageHandler)
	http.HandleFunc("/del", delMessageHandler)
	http.HandleFunc("/all", AllMessageHandler)
	http.HandleFunc("/find", findMessageHandler)

	fmt.Println("Сервер запущен!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("fail programm:", err.Error())
	}
}
