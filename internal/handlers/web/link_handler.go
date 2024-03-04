package web

import (
	"log"
	"net/http"

	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"redirectify/internal/utils"
	"redirectify/internal/views"
)

func MainWebHandler(w http.ResponseWriter, r *http.Request) {
	links, err := models.LinkReadAll(database.Db)

	if err != nil {
		log.Printf("MainWebHandler: %v", err)
	}

	component := views.Links(links)
	err = component.Render(r.Context(), w)

	if err != nil {
		log.Printf("MainWebHandler: %v", err)
		http.Error(w, "Could not render component", http.StatusInternalServerError)
	}
}
