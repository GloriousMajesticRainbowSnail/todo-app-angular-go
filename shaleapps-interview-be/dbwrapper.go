package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DbWrapper provides a simplified interface for querying/updating TODOs
type DbWrapper struct {
	username   string
	password   string
	ip         string
	port       string
	database   string
	connection *sql.DB
}

// NewDbWrapper is the factory method (constructor) for DbWrapper
func NewDbWrapper() *DbWrapper {
	dw := new(DbWrapper)
	dw.username = "root"
	dw.password = "magicbus"
	dw.ip = "127.0.0.1"
	dw.port = "3306"
	dw.database = "TodoDb"
	dw.connection = nil
	return dw
}

// Connect establishes the database connection - must be called prior to use
func (dw *DbWrapper) Connect() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dw.username, dw.password, dw.ip, dw.port, dw.database)
	fmt.Println(connString)
	var err error
	dw.connection, err = sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("Failed to connect to db")
		panic(err.Error())
	}
}

// Close closes the database connection
func (dw *DbWrapper) Close() {
	dw.connection.Close()
}

// GetAllTodos retrieves the full list of todos
func (dw DbWrapper) GetAllTodos() []Todo {
	fmt.Println("GetAllTodos reached")
	var todosList []Todo

	// SELECT Id, TaskDescription, IsComplete FROM Todos;
	results, err := dw.connection.Query("SELECT Id, TaskDescription, IsComplete FROM Todos")
	if err != nil {
		fmt.Println("Failed to retrieve todos")
		panic(err.Error()) // can't really do much if this fails
	}

	fmt.Println("starting results iteration")

	for results.Next() {
		var todo Todo
		err = results.Scan(&todo.ID, &todo.Description, &todo.IsComplete)
		if err != nil {
			panic(err.Error()) // can't really do much if this fails
		}

		todosList = append(todosList, todo)
	}
	return todosList
}
