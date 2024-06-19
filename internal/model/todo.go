package model

import "github.com/gofrs/uuid/v5"

type Todo struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
}

type Todos []Todo

type TodosService interface {
	List() (Todos, error)
	Create(Todo) error
}
