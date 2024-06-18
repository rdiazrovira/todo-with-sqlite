package model

import "github.com/gofrs/uuid/v5"

type Todo struct {
	ID          uuid.UUID `json:"-" db:"id"`
	Title       string    `json:"-" db:"title"`
	Description string    `json:"-" db:"description"`
}

type TodosService interface {
	First() (Todo, error)
}
