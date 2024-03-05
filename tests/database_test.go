package tests

import (
	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/joho/godotenv/autoload"
)

func TestDatabase(t *testing.T) {
	database.New()

	t.Run("Consultar usuários do banco de dados", func(t *testing.T) {
		_, err := models.UsuarioReadAll(database.Db)

		assert.NoError(t, err)
	})
}
