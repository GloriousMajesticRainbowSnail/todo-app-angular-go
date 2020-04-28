-- Create database
DROP DATABASE IF EXISTS TodoDb;
CREATE DATABASE TodoDb;

-- Create table for todos app
USE TodoDb;
CREATE TABLE Todos (
	Id INT NOT NULL AUTO_INCREMENT,
    TaskDescription VARCHAR(1000) NOT NULL,
    IsComplete BOOL NOT NULL,
    PRIMARY KEY (Id)
);
INSERT INTO Todos (
    TaskDescription,
    IsComplete
) VALUES
( 'Initialize the Go server', TRUE ),
( 'Initialize the Angular SPA', TRUE ),
( 'Serve the SPA from the server', TRUE ),
( 'Initialize the database', TRUE ),
( 'Connect to the database from the server', FALSE ),
( 'Create a dedicated database user for the application', FALSE ),
( 'Implement the first end-to-end feature', FALSE ),
( 'Implement the remaining functionality', FALSE ),
( 'Set up Docker', FALSE ),
( 'Deploy solution', FALSE ),
( 'Submit assignment', FALSE );

-- Create user for todos app
DROP USER IF EXISTS 'TodoUser'@'localhost';
CREATE USER 'TodoUser'@'localhost' IDENTIFIED BY '67Gh4cXtydkfL%$';
GRANT INSERT, SELECT, UPDATE ON TodoDb.Todos TO 'TodoUser'@'localhost';