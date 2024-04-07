package tests

import (
	"net/http"
	"net/http/httptest"
	"redirectfy/internal/models"
	"redirectfy/internal/server"
	"redirectfy/internal/services/database"
	"strings"
	"testing"
	
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	// Initialize your mock database and other necessary setup here
	db := database.New()
	defer db.Close()
	
	s := server.Server{ //setup necessário no inicio de cada testsuite dos handlers
		PlanoDeAssinaturaModel: &models.PlanoDeAssinaturaModel{DB: db},
		UsuarioModel:           &models.UsuarioModel{DB: db},
		UsuarioTemporarioModel: &models.UsuarioTemporarioModel{DB: db},
		EmailAutenticacaoModel: &models.EmailAutenticacaoModel{DB: db},
		LinkModel:              &models.LinkModel{DB: db},
		RedirecionadorModel:    &models.RedirecionadorModel{DB: db},
		HistoricoModel:         &models.HistoricoModel{DB: db},
	}
	
	t.Run("Criar usuário handler Successfull", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{ "cpf": 53076281291, "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.NoError(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, http.StatusCreated, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	t.Run("Criar usuário handler FAIL by missing cpf", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": , "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, http.StatusOK, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	t.Run("Criar usuário handler FAIL by cpf wrong", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": 5307628129, "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})

	t.Run("Criar usuário handler FAIL by cpf wrong2", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	t.Run("Criar usuário handler FAIL by missing nome ", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": , "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	t.Run("Criar usuário handler FAIL by nome wrong", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": Guilherme Bernardo, "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})

	t.Run("Criar usuário handler FAIL by missing email", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": "Guilherme Bernardo", "email": , "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	t.Run("Criar usuário handler FAIL by email wrong", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": "Gratuito"}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})

	t.Run("Criar usuário handler FAIL by missing products", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": Gratuito}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})

	t.Run("Criar usuário handler FAIL by products wrong", func(t *testing.T) {
		// Create a new Echo instance
		e := echo.New()
		
		// Define the request body with the correct structure
		JSON := `{"cpf": "53076281291", "nome": "Guilherme Bernardo", "email": "bguilherme51@gmail.com", "products": [{"Name": Gratuito}]}`
		
		// Create a new request
		req := httptest.NewRequest(http.MethodPost, "/api/usuarios_temporarios", strings.NewReader(JSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		// Create a ResponseRecorder to record the response
		rec := httptest.NewRecorder()
		
		// Create a new Context
		c := e.NewContext(req, rec)
		
		if assert.Error(t, s.UsuarioTemporarioCreate(c)) {
			// Call your handler function
			assert.Equal(t, 400, rec.Code)
			// Add more assertions here to check the response body, headers, etc.
		}
	})
	
}
