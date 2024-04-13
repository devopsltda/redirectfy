package utils

import (
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/alexedwards/argon2id"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

var SenhaParams = &argon2id.Params{
	Memory:      19 * 1024,
	Iterations:  2,
	Parallelism: 1,
	SaltLength:  16,
	KeyLength:   32,
}

/*** Variáveis de Ambient ***/
var (
	AppEnv = os.Getenv("APP_ENV")
	Pepper = os.Getenv("PEPPER")
	KirvanoToken = os.Getenv("KIRVANO_TOKEN")
	TempoExpiracao, _ = strconv.Atoi(os.Getenv("VALIDATION_EXPIRE_TIME"))
)

/*** Código Hash ***/
var seededRand *rand.Rand = rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Add(10*time.Second).Unix())))
var symbols []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

func GeraHashCode(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = symbols[seededRand.IntN(len(symbols))]
	}

	return string(b)
}

/*** Validação ***/
var Validate = validator.New()

func IsURLSafe(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' && c != '-' {
			return false
		}
	}

	return true
}

func DebugLog(nomeFuncao, mensagem string, erro error) {
	slog.Debug(nomeFuncao, slog.String("message", mensagem), slog.Any("error", erro))
}

func ErroLog(nomeFuncao, mensagem string, erro error) {
	slog.Error(nomeFuncao, slog.String("message", mensagem), slog.Any("error", erro))
}

func Erro(code int, message string) *echo.HTTPError {
	return echo.NewHTTPError(
		code,
		map[string]string{
			"erros": message,
		},
	)
}

func ErroValidacaoParametro(mensagens []string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		map[string][]string{
			"erros": mensagens,
		},
	)
}
