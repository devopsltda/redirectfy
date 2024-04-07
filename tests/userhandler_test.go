package tests

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"

// 	"redirectfy/internal/handlers/v1/api"
// 	"redirectfy/internal/services/database"
// 	"testing"
// )

// var (
// 	mockDB = map[string]*User{
// 		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
// 	}
// 	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
// )

// func TestUserHandler(t *testing.T) {
// 	database.New()
// 	defer database.Db.Close()
// }
