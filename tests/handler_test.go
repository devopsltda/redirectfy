package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheDevOpsCorp/redirect-max/internal/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHistoricoReadAll(t *testing.T) {
	e := echo.New()
	s := server.NewTestServer()

	req := httptest.NewRequest(http.MethodGet, "/api/historico", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, s.HistoricoReadAll(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}
