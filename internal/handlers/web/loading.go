package web

import (
	"log"
	"net/http"

	"redirectify/internal/views"
)

func LoadingWebHandler(w http.ResponseWriter, r *http.Request) {
	component := views.Loading()
	err := component.Render(r.Context(), w)

	if err != nil {
		log.Printf("LoadingWebHandler: %v", err)
		http.Error(w, "A página não pôde ser renderizada.", http.StatusInternalServerError)
	}
}
