package main

import (
	"github.com/google/uuid"
)

type player struct {
	id uuid.UUID
}

func (player) load(id uuid.UUID) player {

	p := player{id: id}

	return p
}

type inventory struct {
	parts []string
}

func main() {

}
