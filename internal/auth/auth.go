package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"redirectify/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ChaveDeAcesso  = os.Getenv("JWT_SECRET")
	ChaveDeRefresh = os.Getenv("JWT_REFRESH_SECRET")
)

type Claims struct {
	Id            int64  `json:"id"`
	Nome          string `json:"nome"`
	NomeDeUsuario string `json:"nome_de_usuario"`
	Autenticado   bool   `json:"autenticado"`
	jwt.RegisteredClaims
}

func GeraTokensESetaCookies(id int64, nome, nomeDeUsuario string, autenticado bool, c echo.Context) error {
	accessToken, exp, err := GeraTokenAcesso(id, nome, nomeDeUsuario, autenticado)

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)
	SetCookieUsuario(nomeDeUsuario, exp, c)

	refreshToken, exp, err := GeraTokenRefresh(id, nome, nomeDeUsuario, autenticado)

	if err != nil {
		return err
	}

	SetCookieToken("refresh-token", refreshToken, exp, c)

	return nil
}

func GeraTokenAcesso(id int64, nome, nomeDeUsuario string, autenticado bool) (string, time.Time, error) {
	expiraEm := time.Now().Add(1 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, autenticado, expiraEm, []byte(ChaveDeAcesso))
}

func GeraTokenRefresh(id int64, nome, nomeDeUsuario string, autenticado bool) (string, time.Time, error) {
	expiraEm := time.Now().Add(24 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, autenticado, expiraEm, []byte(ChaveDeRefresh))
}

func GeraToken(id int64, nome, nomeDeUsuario string, autenticado bool, expiraEm time.Time, chave []byte) (string, time.Time, error) {
	if !autenticado {
		return "", time.Now(), fmt.Errorf(utils.MensagemUsuarioNaoAutenticado)
	}

	claims := &Claims{
		Id:            id,
		Nome:          nome,
		NomeDeUsuario: nomeDeUsuario,
		Autenticado:   autenticado,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expiraEm},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(chave)

	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expiraEm, nil
}

func VerificaToken(tokenFornecido string) error {
	token, err := jwt.Parse(tokenFornecido, func(t *jwt.Token) (interface{}, error) {
		return ChaveDeAcesso, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf(utils.MensagemJWTInvalido)
	}

	return nil
}

func SetCookieToken(nome, token string, expiraEm time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = nome
	cookie.Value = token
	cookie.Expires = expiraEm
	cookie.Path = "/"

	c.SetCookie(cookie)
}

func SetCookieUsuario(nomeDeUsuario string, expiraEm time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "usuario"
	cookie.Value = nomeDeUsuario
	cookie.Expires = expiraEm
	cookie.Path = "/"

	c.SetCookie(cookie)
}

func PathWithNoAuthRequired(c echo.Context) bool {
	return (c.Request().URL.Path == "/v1/api/usuarios/login" && c.Request().Method == "POST") ||
		(c.Request().URL.Path == "/v1/api/usuarios" && c.Request().Method == "POST") ||
		(c.Path() == "/v1/api/usuarios/troca_de_senha/:valor" && c.Request().Method == "PATCH") ||
		(c.Path() == "/v1/api/usuarios/:nome_de_usuario/troca_de_senha" && c.Request().Method == "POST") ||
		(c.Path() == "/v1/api/docs/*" && c.Request().Method == "GET") ||
		(c.Path() == "/v1/api/autenticacao/:valor" && c.Request().Method == "PATCH")
}

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("usuario") == nil || PathWithNoAuthRequired(c) {
			return next(c)
		}

		nomeDeUsuario := c.Get("usuario").(*jwt.Token)

		claims := nomeDeUsuario.Claims.(*Claims)

		if !claims.Autenticado {
			return utils.ErroUsuarioNaoAutenticado
		}

		if time.Until(claims.RegisteredClaims.ExpiresAt.Time) < 15*time.Minute {
			refreshCookie, err := c.Cookie("refresh-token")

			if err == nil && refreshCookie != nil {
				token, err := jwt.ParseWithClaims(refreshCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(ChaveDeRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						return utils.ErroAssinaturaJWT
					}
				}

				if token != nil && token.Valid {
					err = GeraTokensESetaCookies(claims.Id, claims.Nome, claims.NomeDeUsuario, claims.Autenticado, c)

					if err != nil {
						return utils.ErroAssinaturaJWT
					}
				}
			}
		}

		return next(c)
	}
}
