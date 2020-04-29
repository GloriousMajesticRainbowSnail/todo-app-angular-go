# todo-app-angular-go
A quick-and-dirty todo app written in Angular and Go

## Set-up instructions

### Front-end

This project uses [Angular CLI](https://cli.angular.io/). To build the production bundle, which is currently configured to output to the `fe-bundle` directory, run `ng build --prod` from the `todo-fe` directory. The bundle will be served from the server.

### Server

The Go files are contained in the `todo-be` directory and can be run with the command `go run ./todo-be`. 

### Database

The included SQL scripts are written for MySQL. Execute `db-scripts\database_init.sql` to initialize the database, table, and user creation.
