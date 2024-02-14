package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheDevOpsCorp/redirect-max/internal/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLinkReadByCodigoHash(t *testing.T)  {
	e := echo.New()
	s := server.NewTestServer()
}

func TestLinkReadAll(t *testing.T) {
	e := echo.New()
	s := server.NewTestServer()
}

func TestLinkCreate(t *testing.T) {
	e := echo.New()
	s := server.NewTestServer()
}

func TestLinkUpdate(t *testing.T) {
	e := echo.New()
	s := server.NewTestServer()
}

func TestLinkRemove(t *testing.T) {
	e := echo.New()
	s := server.NewTestServer()
}