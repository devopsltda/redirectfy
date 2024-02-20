package tests

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/TheDevOpsCorp/redirectify/internal/auth"
// 	"github.com/TheDevOpsCorp/redirectify/internal/handlers/v1/api"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/golang-jwt/jwt/v5"
// )

// // TestSuiteHistoricoReadAll contains all tests for the HistoricoReadAll handler
// func TestHistoricoUsuarioReadAll(t *testing.T) {
// 	// Create a new Echo instance
// 	e := echo.New()

// 	// Create a valid JWT token
// 	claims := jwt.MapClaims{
// 		"username": "testuser",
// 		// Add other claims as necessary
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token
// 	signedToken, _ := token.SignedString([]byte(auth.ChaveDeAcesso))

// 	// Create a new request with the JWT token in the cookie
// 	req := httptest.NewRequest(http.MethodGet, "/api/usuarios/historico", nil)
// 	req.AddCookie(&http.Cookie{
// 		Name:  "access-token",
// 		Value: signedToken,
// 	})

// 	// Create a response recorder
// 	res := httptest.NewRecorder()

// 	// Create a new Echo context
// 	c := e.NewContext(req, res)

// 	// Call the handler
// 	err := api.HistoricoUsuarioReadAll(c)

// 	// Assert the expected response
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, res.Code)
// 	// Additional assertions for the response body
// }