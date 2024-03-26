package tests

import (
	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	database.New()
	defer database.Db.Close()

	t.Run("Adicionar usuário ao banco de dados", func(t *testing.T) {
		_, err := models.UsuarioCreate(
			database.Db,
			"02958985261",
			"Eduardo Henrique Freire Machado",
			"edu_hen_fm",
			"edu.hen.fm@gmail.com",
			"senha-muito-complexa-aqui",
			"2001-08-30",
			1,
		)

		assert.NoError(t, err)
	})

	t.Run("Consultar usuários do banco de dados", func(t *testing.T) {
		_, err := models.UsuarioReadAll(database.Db)

		assert.NoError(t, err)
	})
}
