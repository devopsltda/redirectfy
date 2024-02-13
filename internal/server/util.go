package server

import (
	"net/http"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

/*** Validação ***/
var Validate = validator.New()

func validaNomeDeUsuario(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_' && c != '-' {
			return false
		}
	}

	return true
}

/*** Erro ***/
type Erro echo.HTTPError // @name Erro

var ErroSenha *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "A senha fornecida está incorreta.")
var ErroCriacaoSenha *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "Ocorreu um erro ao criar a senha.")
var ErroUsuario *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "O nome de usuário ou email está incorreto.")
var ErroAssinaturaJWT *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "Ocorreu um erro na assinatura do token JWT.")
var ErroBancoDados *echo.HTTPError = echo.NewHTTPError(http.StatusInternalServerError, "Ocorreu um erro no banco de dados.")
var ErroValidacaoNome *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "Por favor, forneça um nome válido.")
var ErroValidacaoCodigoHash *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "Por favor, forneça um código hash válido (apenas contém letras, números ou os símbolos '-' e '_' e tem 10 caracteres).")
var ErroValidacaoNomeDeUsuario *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "Por favor, forneça um nome de usuário válido (apenas contém letras, números ou os símbolos '-' e '_').")

func ErroValidacaoParametro(mensagem []string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		map[string][]string{
			"Erros": mensagem,
		},
	)
}
