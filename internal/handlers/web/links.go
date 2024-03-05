package web

import (
	"database/sql"
	"log"
	"net/http"

	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"redirectify/internal/utils"
	"redirectify/internal/views"

	"github.com/labstack/echo/v4"
)

func GoToLinkWebHandler(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Insira o HTML de erro 4xx aqui")
	}

	links, err := models.LinkReadAll(database.Db, codigoHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Insira o HTML de erro 404 aqui")
		} else {
			log.Printf("GoToLinkWebHandler: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Insira o HTML de erro 5xx aqui")
		}
	}

	linkWhatsapp, linkTelegram, err := models.LinkPicker(links)

	component := views.GoToLink(linkWhatsapp, linkTelegram)
	err = component.Render(c.Request().Context(), c.Response().Writer)

	if err != nil {
		log.Printf("GoToLinkWebHandler: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Insira o HTML de erro 5xx aqui")
	}

	return nil
}
