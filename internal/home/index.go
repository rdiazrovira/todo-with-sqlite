package home

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting...")
	http.Redirect(w, r, "/todos", http.StatusFound)
}
