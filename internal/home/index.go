package home

import (
	"net/http"
	"todox-with-sqlite/internal/model"

	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	todos := r.Context().Value("todos").(model.TodosService)

	todo, err := todos.First()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rw.Set("todo", todo)
	if err := rw.Render("home/index.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
