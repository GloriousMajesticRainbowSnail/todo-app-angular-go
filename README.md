# todo-app-angular-go
A quick-and-dirty todo app written in Angular, Go, and MySQL.

## Set-up instructions

### Front-end

The front end is built with [Angular CLI](https://cli.angular.io/). To build the production bundle, which is currently configured to output to the `fe-bundle` directory, run `ng build --prod` from the `todo-fe` directory. *The bundle will be served from the server, which expects the bundle to be located at `fe-bundle`.* 

(This has been pre-compiled and included in the repo due to memory requirements for using the `ng build --prod` command - see [here](https://github.com/angular/angular-cli/issues/10529) for details.)

Styling is off-the-shelf Bootstrap - there is a styling issue (see below), but fixing this would take additional investigation and might require configuring or overriding Bootstrap.

### Server

The Go files are contained in the `todo-be` directory and can be run with the command `go run ./todo-be`, or by running `go build ./todo-be` and executing the resulting binary. You will need to install the mysql driver - instructions [here](https://github.com/Go-SQL-Driver/MySQL/). 

The server is hard-coded to serve from port `8081`.

### Database

The included SQL scripts are written for MySQL and assume root-level access to a MySQL instance serving the default port `3306`. Execute `db-scripts\database_init.sql` to initialize the database, table, and user creation. Database connection details are hard-coded in the server code.

## Improvements/cut corners

- Authentication/users are not implemented; everyone can see the same set of TODOs.
- Configuration for database connection et. al. has not been implemented.
- No unit-testing; this was implemented as a learning/exploratory prototype. I'm otherwise generally in favour of test-driven development.
- No transactions have been implemented for the database queries, which is a potential source of bugs and data loss. I didn't take the time to implement this because this is a quick-and-dirty learning project.
- The Go code is generally messy and fairly disorganized; I would like to spend some time refactoring this and aligning with best practices for Go, which I'm still learning.
- Todo description text breaks out of the styling container if it's too long/if the container is too narrow. This is especially obvious on mobile. I haven't looked into this yet, but it's default Bootstrap behaviour so I assume configuration/overriding would be required.
- As noted in the `server.go` file, the REST API implementation is hacky and falls short of best-practice for REST API design:
``` golang
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
```

## Feedback and usage

Feel free to provide comments, open up issues, etc. Any advice or suggestions are welcome. 

I would not recommend copying this solution for your own learning unless you are an experienced developer and understand the noted cut corners above. I will clean this up over time so that, hopefully, it will provide a better frame-of-reference for people learning Angular and/or Go. If you have a question, feel free to comment, open an issue, etc.
