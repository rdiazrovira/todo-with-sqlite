package todos

import (
	"todox-with-sqlite/internal/model"

	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) model.TodosService {
	return &service{db}
}

func (s *service) List() (todos model.Todos, err error) {
	return todos, s.db.Select(&todos, "SELECT * FROM todos")
}

func (s *service) Create(todo model.Todo) error {
	todo.ID = uuid.Must(uuid.NewV4())
	todo.Description = "N/A"
	_, err := s.db.NamedExec(`INSERT INTO todos (id, title, description) VALUES (:id, :title, :description)`, todo)
	return err
}
