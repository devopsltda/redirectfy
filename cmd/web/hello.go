package web

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/cmd/web/views"
)

func LoginWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Printf("LoginWebHandler: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := views.HelloPost(name)
	component.Render(r.Context(), w)
}
