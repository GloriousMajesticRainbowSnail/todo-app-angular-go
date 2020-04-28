package main

// Todo is the model for a to-do item
type Todo struct {
	ID          int64  `json:"Id"`
	Description string `json:"Description"`
	IsComplete  bool   `json:"IsComplete"`
}
