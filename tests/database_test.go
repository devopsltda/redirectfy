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

	t.Run("Consultar usuários do banco de dados", func(t *testing.T) {
		_, err := models.UsuarioReadAll(database.Db)

		assert.NoError(t, err)
	})
}
