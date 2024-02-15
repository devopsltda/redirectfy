package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheDevOpsCorp/redirectify/internal/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestSuiteHistoricoReadAll contains all tests for the HistoricoReadAll handler
func TestSuiteHistoricoReadAll(t *testing.T) {
	// Initialize the test server
	s := server.NewTestServer()
	e := echo.New()

	// Define the common request setup
	setupRequest := func(method, url string) (*httptest.ResponseRecorder, *echo.Context) {
		req := httptest.NewRequest(method, url, nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		return res, &c
	}

	// Test cases
	t.Run("Pass 200: successful history retrieval", func(t *testing.T) {
		res, c := setupRequest(http.MethodGet, "/api/historico")

		err := s.HistoricoReadAll(*c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}
