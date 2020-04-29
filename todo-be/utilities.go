package main

import "log"

// HandleError describes error handling
func HandleError(e error) {
	if e != nil {
		// What's the best practice for handling errors in Go?
		log.Fatal(e.Error())
		panic(e.Error())
	}
}
