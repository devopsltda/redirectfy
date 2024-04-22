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

// Esses parâmetros são os mínimos recomendados para
// o uso do algoritmo argon2 de forma eficiente:
// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
var SenhaParams = &argon2id.Params{
	Memory:      19 * 1024,
	Iterations:  2,
	Parallelism: 1,
	SaltLength:  16,
	KeyLength:   32,
}

/*** Variáveis de Ambiente ***/
var (
	AppEnv = os.Getenv("APP_ENV")
	Pepper = os.Getenv("PEPPER")
	KirvanoToken = os.Getenv("KIRVANO_TOKEN")
	TempoExpiracao, _ = strconv.Atoi(os.Getenv("VALIDATION_EXPIRE_TIME"))
)

/*** Código Hash ***/
var seededRand *rand.Rand = rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Add(10*time.Second).Unix())))
var symbols []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

// Gera o hash selecionando de forma aleatória caracteres alfanuméricos, '_' ou '-'
// no tamanho inserido.
func GeraHashCode(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = symbols[seededRand.IntN(len(symbols))]
	}

	return string(b)
}

/*** Validação ***/
var Validate = validator.New()

// Verifica se s contém algum caractere não alfanumérico que seja diferente
// de '_' ou '-'.
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

// Volta a mensagem de erro padronizada.
func Erro(code int, message string) *echo.HTTPError {
	return echo.NewHTTPError(
		code,
		[]string{message},
	)
}

func ErroValidacaoParametro(mensagens []string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		mensagens,
	)
}
