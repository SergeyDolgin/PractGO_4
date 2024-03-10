// Что нужно сделать:
// Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом: **
// 1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу.
// Данный запрос должен возвращать ID пользователя и статус 201.
// 2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья.
// 3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")
	r.HandleFunc("/", handleGet).Methods("GET")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()

}

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "post\n")
}

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}

type Activities struct {
	activities []Activity
}

func (c *Activities) Insert(activity Activity) uint64 {
	activity.ID = uint64(len(c.activities))
	c.activities = append(c.activities, activity)
	return activity.ID
}

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	var ErrIDNotFound = fmt.Errorf("ID not found")
	if id >= uint64(len(c.activities)) {
		return Activity{}, ErrIDNotFound
	}
	return c.activities[id], nil
}

type httpServer struct {
	Activities *Activities
}

type IDDocument struct {
	ID uint64 `json:"id"`
}

func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		Activities: &Activities{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/", server.handlePost).Methods("POST")
	r.HandleFunc("/", server.handleGet).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
