package main

import (
	"database/sql"
	"encoding/json" // used to serialize/deserialize json req/res
	"fmt"           // used to write output
	"log"           // logging
	"net/http"      // used to handle http req/res
	"strconv"       // used for converting primitives to string

	_ "github.com/go-sql-driver/mysql"
)

// Todo is the model for a to-do item
type Todo struct {
	ID          int    `json:"Id"`
	Description string `json:"Description"`
	IsComplete  bool   `json:"IsComplete"`
}

// Todos is a list of dummy todos (until database is connected)
var Todos []Todo

var folderDist = "fe-bundle"

func initDummyData() {
	Todos = []Todo{
		Todo{ID: 1, Description: "Item 1", IsComplete: false},
		Todo{ID: 2, Description: "Item 2", IsComplete: true},
	}
}

func parseTodo(r *http.Request) Todo {
	var body Todo
	json.NewDecoder(r.Body).Decode(&body)
	return body
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URI: " + r.RequestURI + ", method: " + r.Method)
	/*
		Best practice for REST API is to use the format "/todos/{id}", where:
		- GET /todos returns all todos
		- GET /todos/{id} returns a specific todo
		- POST /todos/{id} creates a new todo
		- PUT /todos/{id} updates a specific todo
		- etc.

		Due to the relative complexity of setting up routing in Go
		(the default HTTP package does not include an intuitive REST-endpoint-friendly API,
		and the commonly-used Gorilla MUX router doesn't play nice with serving an Angular
		SPA without configuration/customization - most online tutorials direct the developer
		to serve the SPA from a separate server), I have opted to use the HTTP/REST method
		to determine which operation should be performed/at which level. This is a time-saving
		decision made in part because this is a quick-and-dirty interview app, and in part because
		I'm learning Go and Docker for this assignment and thus have to limit scope.

		I know there must be a conventional/practical way to implement "correct" REST behaviour,
		and I look forward to learning that either from experienced Go developers or on my own when
		I've had more time to learn!

		Also, I apologize for the multi-line comment-novel; this is normally something I would discuss
		with the team instead of documenting in a long comment, but I don't know whether or not I will
		have the opportunity to do so. I would love	to take some time to "code-review" this solution
		and learn how my code could be improved!
	*/
	switch r.Method {
	case "GET":
		// Assume this is retrieving all TODOs
		json.NewEncoder(w).Encode(Todos)
		break
	case "PUT":
		// Assume this is updating an individual TODO
		var b = parseTodo(r)
		fmt.Println("updated " + strconv.Itoa(b.ID))
		json.NewEncoder(w).Encode(b)
		break
	case "POST":
		// Assume this is creating a new TODO
		var b = parseTodo(r)
		fmt.Println("created todo for " + b.Description)
		json.NewEncoder(w).Encode(b)
	}
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir(folderDist)))
	http.HandleFunc("/todos", handleTodos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func connectDb() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func main() {
	initDummyData()
	handleRequests()
	connectDb()
	fmt.Println("Server started!")
}
