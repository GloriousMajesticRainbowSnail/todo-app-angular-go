package main

import (
	"encoding/json" // used to serialize/deserialize json req/res
	"fmt"           // used to write output
	"log"           // logging
	"net/http"      // used to handle http req/res
)

// Todo is the model for a to-do item
type Todo struct {
	Name       string `json:"Name"`
	IsComplete bool   `json:"IsComplete"`
}

// Todos is a list of dummy todos (until database is connected)
var Todos []Todo

func initDummyData() {
	Todos = []Todo{
		Todo{Name: "Item 1", IsComplete: false},
		Todo{Name: "Item 2", IsComplete: true},
	}
}

func printEndpoint(r *http.Request, msg string) {
	fmt.Println("Endpoint hit: " + msg + ", request URI: " + r.RequestURI)
}

func showDefault(w http.ResponseWriter, r *http.Request) {
	printEndpoint(r, "showDefault") // prints twice - one for /, one for favicon
	fmt.Fprintf(w, "The server is running.")
}

func returnAllTodos(w http.ResponseWriter, r *http.Request) {
	printEndpoint(r, "returnAllTodos")
	json.NewEncoder(w).Encode(Todos)
}

func handleRequests() {
	http.HandleFunc("/", showDefault)
	http.HandleFunc("/todos", returnAllTodos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	initDummyData()
	handleRequests()
	fmt.Println("Server started!")
}
