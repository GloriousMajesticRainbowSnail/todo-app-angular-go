package main

import (
	"encoding/json" // used to serialize/deserialize json req/res
	"errors"
	"fmt"      // used to write output
	"log"      // logging
	"net/http" // used to handle http req/res
	"strconv"
	// used for converting primitives to string
)

const folderDist = "../fe-bundle"

func parseTodo(r *http.Request) Todo {
	var body Todo
	json.NewDecoder(r.Body).Decode(&body)
	return body
}

func errorHandler(w http.ResponseWriter, statusCode int, err error, resMsg string) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, resMsg)
	fmt.Println(err.Error())
}

func createTodosHandler() (func(w http.ResponseWriter, r *http.Request), func()) {
	db := NewDbWrapper()
	db.Connect()
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request URI: " + r.RequestURI + ", method: " + r.Method)
		// See comment at end-of-file
		switch r.Method {
		case "GET":
			// Assume "GET /todos" is retrieving all TODOs
			todos := db.GetAllTodos()
			json.NewEncoder(w).Encode(todos)
		case "PUT":
			// Assume "PUT /todos" is updating the provided TODO
			var updated = parseTodo(r)
			updated = db.UpdateTodo(updated)
			json.NewEncoder(w).Encode(updated)
			break
		case "POST":
			// Assume "POST /todos" is creating the provided TODO
			var newTodo = parseTodo(r)
			newTodo = db.CreateTodo(newTodo)
			json.NewEncoder(w).Encode(newTodo)
			break
		case "DELETE":
			// Use query parameter to identify the TODO to delete
			idString := r.URL.Query().Get("id")
			if idString == "" {
				errorHandler(w, http.StatusBadRequest, errors.New("Invalid ID string: "+idString), "Unable to delete: no ID provided")
			}

			id, err := strconv.Atoi(idString)
			if id <= 0 || err != nil {
				errorHandler(w, http.StatusBadRequest, err, "Unable to delete: invalid ID")
			}

			id, err = db.RemoveTodo(id)
			if err != nil {
				errorHandler(w, http.StatusInternalServerError, err, "Failed to delete id: "+idString)
			}

			json.NewEncoder(w).Encode(id)
		}
	}, db.Close
}

func main() {
	// Serve the Angular app
	http.Handle("/", http.FileServer(http.Dir(folderDist)))

	// Handle the REST API endpoint "todos"
	todosHandler, dbCloser := createTodosHandler()
	http.HandleFunc("/todos", todosHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8081", nil))
	fmt.Println("Server started!")

	// Close the db connection when finished
	defer dbCloser()
}

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
