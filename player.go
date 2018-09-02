package main

import "fmt"

type player struct {
	ID   *string `json:"_id" bson:"_id"`
	Name *string `json:"name"`
}

func (p player) String() string {
	return fmt.Sprint("[Player| id:", p.ID, ", name:", p.Name, "]")
}
