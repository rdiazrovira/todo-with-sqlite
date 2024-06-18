package todos

import (
	"todox-with-sqlite/internal/model"

	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) model.TodosService {
	return &service{db}
}

func (s *service) First() (model.Todo, error) {
	todo := model.Todo{}
	return todo, s.db.Get(&todo, "SELECT * FROM todos LIMIT 1")
}
