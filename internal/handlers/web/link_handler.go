package web

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirectify/internal/components"
	"github.com/TheDevOpsCorp/redirectify/internal/database"
	"github.com/TheDevOpsCorp/redirectify/internal/models"
	"github.com/TheDevOpsCorp/redirectify/internal/utils"
	"github.com/TheDevOpsCorp/redirectify/internal/views"
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

func LinkAccessWebHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidaNomeDeUsuario(r.URL.Path[1:]) {
		http.Error(w, "Não foi Encontrado", http.StatusNotFound)
		return
	}

	link, err := models.LinkReadByCodigoHash(database.Db, r.URL.Path[1:])

	if err != nil {
		log.Printf("LinkAccessWebHandler: %v", err)
	}

	component := views.LinkAccess(link)
	err = component.Render(r.Context(), w)

	if err != nil {
		log.Printf("LinkAccessWebHandler: %v", err)
		http.Error(w, "Could not render component", http.StatusInternalServerError)
	}
}

func LinkCreateWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Printf("LinkCreateWebHandler: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	nome := r.FormValue("link_create_nome")
	linkWhatsapp := r.FormValue("link_create_link_whatsapp")
	linkTelegram := r.FormValue("link_create_link_telegram")
	ordemDeRedirecionamento := r.FormValue("link_create_ordem_de_redirecionamento")

	var codigoHash string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHash = utils.GeraHashCode(10)

		codigoHashExiste, err = models.LinkCheckIfCodigoHashExists(database.Db, codigoHash)

		if err != nil {
			log.Printf("LinkCreateWebHandler: %v", err)
		}
	}

	err = models.LinkCreate(
		database.Db,
		nome,
		codigoHash,
		linkWhatsapp,
		linkTelegram,
		ordemDeRedirecionamento,
		1,
	)

	if err != nil {
		log.Printf("LinkCreateWebHandler: %v", err)
	}

	component := components.LinkCreateSuccessful()
	err = component.Render(r.Context(), w)

	if err != nil {
		log.Printf("LinkCreateWebHandler: %v", err)
		http.Error(w, "Could not render component", http.StatusInternalServerError)
	}
}

func LinkCreateFormWebHandler(w http.ResponseWriter, r *http.Request) {
	component := components.LinkCreateForm()
	err := component.Render(r.Context(), w)

	if err != nil {
		log.Printf("LinkCreateFormWebHandler: %v", err)
		http.Error(w, "Could not render component", http.StatusInternalServerError)
	}
}

func LinkCreateButtonWebHandler(w http.ResponseWriter, r *http.Request) {
	component := components.LinkCreateButton()
	err := component.Render(r.Context(), w)

	if err != nil {
		log.Printf("LinkCreateButtonWebHandler: %v", err)
		http.Error(w, "Could not render component", http.StatusInternalServerError)
	}
}
