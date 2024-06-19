package todos

import (
	"net/http"
	"todox-with-sqlite/internal/model"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}

	if err := form.Decode(r, &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todos := r.Context().Value("todos").(*service)
	if err := todos.Create(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	if err := rw.RenderClean("todos/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	todos := r.Context().Value("todos").(*service)

	list, err := todos.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("list", list)

	if err := rw.Render("todos/index.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
