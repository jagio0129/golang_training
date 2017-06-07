package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "write persentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo Show: %s", ps.ByName("todoId"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
