package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DbWrapper provides a simplified interface for querying/updating TODOs
type DbWrapper struct {
	username string
	password string
	ip       string
	port     string
	database string
	dbHandle *sql.DB
}

// NewDbWrapper is the factory method (constructor) for DbWrapper
func NewDbWrapper() *DbWrapper {
	dw := new(DbWrapper)
	dw.username = "TodoUser"
	dw.password = "67Gh4cXtydkfL%$"
	dw.ip = "127.0.0.1"
	dw.port = "3306"
	dw.database = "TodoDb"
	return dw
}

// Connect establishes the database connection - must be called prior to use
func (dw *DbWrapper) Connect() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dw.username, dw.password, dw.ip, dw.port, dw.database)
	dbHandle, err := sql.Open("mysql", connString)
	HandleError(err)

	/*
		database/sql doesn't establish a connection until they're needed;
		ping DB to ensure credentials are valid and connection is successful
	*/
	err = dbHandle.Ping()
	HandleError(err)

	dw.dbHandle = dbHandle
}

// Close closes the database connection
func (dw *DbWrapper) Close() {
	dw.dbHandle.Close()
}

// GetAllTodos retrieves the full list of todos
func (dw DbWrapper) GetAllTodos() []Todo {
	var todosList []Todo

	// SELECT Id, TaskDescription, IsComplete FROM Todos;
	results, err := dw.dbHandle.Query("SELECT Id, TaskDescription, IsComplete FROM Todos")
	HandleError(err)

	for results.Next() {
		var todo Todo
		err = results.Scan(&todo.ID, &todo.Description, &todo.IsComplete)
		HandleError(err)

		todosList = append(todosList, todo)
	}

	return todosList
}

// UpdateTodo persists changes to an individual todo
func (dw DbWrapper) UpdateTodo(todo Todo) Todo {
	/*
		UPDATE Todos
		SET IsComplete={Todo.IsComplete}
		WHERE Id = {Todo.ID};
	*/
	stmt, err := dw.dbHandle.Prepare("UPDATE Todos SET IsComplete=? WHERE Id=?")
	HandleError(err)
	results, err := stmt.Exec(todo.IsComplete, todo.ID)
	HandleError(err)
	affected, err := results.RowsAffected()
	HandleError(err)

	if affected < 1 {
		HandleError(errors.New("Update failed: no rows affected"))
	}

	if affected > 1 {
		HandleError(errors.New("Update failed: multiple rows affected"))
	}

	return todo
}

// CreateTodo persists a new todo
func (dw DbWrapper) CreateTodo(todo Todo) Todo {
	/*
		INSERT INTO Todos (
			TaskDescription,
			IsComplete
		) VALUES (
			{Todo.TaskDescription},
			FALSE
		);
	*/
	stmt, err := dw.dbHandle.Prepare("INSERT INTO Todos (TaskDescription, IsComplete) VALUES (?, FALSE)")
	HandleError(err)
	results, err := stmt.Exec(todo.Description)
	HandleError(err)
	inserted, err := results.LastInsertId()
	HandleError(err)

	if inserted < 1 {
		HandleError(errors.New("Insert failed: invalid inserted ID"))
	}

	return todo
}
