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

func ValidaNomeDeUsuario(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' && c != '-' {
			return false
		}
	}

	return true
}

/*** Mensagens de Confirmação ***/
var MensagemPlanoDeAssinaturaCriadoComSucesso = "O plano de assinatura foi adicionado com sucesso."
var MensagemPlanoDeAssinaturaAtualizadoComSucesso = "O plano de assinatura foi atualizado com sucesso."
var MensagemPlanoDeAssinaturaRemovidoComSucesso = "O plano de assinatura foi removido com sucesso."

var MensagemRedirecionadorAtualizadoComSucesso = "O redirecionador foi atualizado com sucesso."
var MensagemRedirecionadorRemovidoComSucesso = "O redirecionador foi removido com sucesso."

var MensagemLinkAtualizadoComSucesso = "O link foi atualizado com sucesso."
var MensagemLinkRemovidoComSucesso = "O link foi removido com sucesso."

var MensagemUsuarioCriadoComSucesso = "O usuário foi adicionado com sucesso."
var MensagemUsuarioAtualizadoComSucesso = "O usuário foi atualizado com sucesso."
var MensagemUsuarioAutenticadoComSucesso = "O usuário foi autenticado com sucesso."
var MensagemUsuarioSenhaTrocadaComSucesso = "A senha do usuário foi trocada com sucesso."
var MensagemUsuarioRemovidoComSucesso = "O usuário foi removido com sucesso."
var MensagemUsuarioLogadoComSucesso = "O usuário foi logado com sucesso."
var MensagemUsuarioNaoAutenticado = "O usuário não está autenticado."

var MensagemJWTInvalido = "Token JWT Inválido."

var ErroLogin = echo.NewHTTPError(http.StatusBadRequest, []string{"O email ou senha fornecidos estão incorretos."})
var ErroCriacaoSenha = echo.NewHTTPError(http.StatusBadRequest, []string{"Ocorreu um erro ao criar a senha."})
var ErroAssinaturaJWT = echo.NewHTTPError(http.StatusBadRequest, []string{"Ocorreu um erro na assinatura do token JWT."})
var ErroBancoDados = echo.NewHTTPError(http.StatusInternalServerError, []string{"Ocorreu um erro no banco de dados."})
var ErroValidacaoNome = echo.NewHTTPError(http.StatusBadRequest, []string{"Por favor, forneça um nome válido."})
var ErroValidacaoCodigoHash = echo.NewHTTPError(http.StatusBadRequest, []string{"Por favor, forneça um código hash válido (apenas contém letras, números ou os símbolos '-' e '_' e tem 10 caracteres)."})
var ErroValidacaoNomeDeUsuario = echo.NewHTTPError(http.StatusBadRequest, []string{"Por favor, forneça um nome de usuário válido (apenas contém letras, números ou os símbolos '-' e '_')."})
var ErroUsuarioNaoAutenticado = echo.NewHTTPError(http.StatusUnauthorized, []string{"Por favor, autentique seu usuário no email enviado ou solicite um novo email."})

func DebugLog(nomeFuncao, mensagem string, erro error) {
	slog.Error(nomeFuncao, slog.String("message", mensagem), slog.Any("error", erro))
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
