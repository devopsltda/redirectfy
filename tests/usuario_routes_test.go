package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheDevOpsCorp/redirectify/internal/handlers/v1/api"
	"github.com/TheDevOpsCorp/redirectify/internal/server"
	"github.com/TheDevOpsCorp/redirectify/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestSuiteUsuario contains all tests for the Usuario handlers
func TestSuiteUsuario(t *testing.T) {
	// Initialize the test server
	_ = server.NewTestServer()
	e := echo.New()

	// Define the common request setup
	setupRequest := func(method, url string) (*httptest.ResponseRecorder, *echo.Context) {
		req := httptest.NewRequest(method, url, nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		return res, &c
	}

	var usuarioReadAllMock = api.UsuarioReadAll

	// Test cases
	t.Run("Pass  200: successful user retrieval", func(t *testing.T) {
		res, c := setupRequest(http.MethodGet, "/api/usuarios")

		err := api.UsuarioReadAll(*c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		expectedBody := `{"message": "expected response body"}` //FALTA ALTERAR AQUI AINDA
		assert.Equal(t, expectedBody, res.Body.String())
	})

	t.Run("Pass  200: successful user retrieval by username", func(t *testing.T) {
		res, c := setupRequest(http.MethodGet, "/api/usuarios/testuser")

		err := api.UsuarioReadByNomeDeUsuario(*c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the response body
		expectedBody := `{"message": "expected response body"}` //FALTA ALTERAR AQUI AINDA
		assert.Equal(t, expectedBody, res.Body.String())
	})
	t.Run("Fail   500: database error", func(t *testing.T) {
		res, c := setupRequest(http.MethodGet, "/api/usuarios")

		// Mock the error
		usuarioReadAllMock = func(c echo.Context) error {
			return utils.ErroBancoDados
		}

		err := usuarioReadAllMock(*c)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
