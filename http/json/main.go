package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

type User struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Age     int     `json:"age"`
	Married bool    `json:"married"`
	Height  float64 `json:"height"`
}

func (u User) Println() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Address:", u.Address)
	fmt.Println("Age:", u.Age)
	fmt.Println("Married:", u.Married)
	fmt.Println("Height:", u.Height)
}

var mtx sync.Mutex

var allUser = make([]User, 0)

type HttpResponse struct {
	Id   int `json:"id"`
	AllU []User
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail method"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mtx.Lock()
	allUser = append(allUser, user)
	mtx.Unlock()

	//user.Println()

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		msg := "fail method"
		fmt.Println(msg)

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write HTTP body:", err.Error())
		}
		return
	}

	httpResponse := HttpResponse{
		Id:   rand.Intn(100),
		AllU: allUser,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("fail to write HTTP body:", err.Error())
		return
	}

}

func main() {
	http.HandleFunc("/post", postUserHandler)
	http.HandleFunc("/get", getUserHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка сервера:", err.Error())
	}

}
