package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Id            int64  `json:"id"`
	Nome          string `json:"nome"`
	NomeDeUsuario string `json:"nome_de_usuario"`
	jwt.RegisteredClaims
}

var ChaveDeAcesso = os.Getenv("JWT_SECRET")
var ChaveDeRefresh = os.Getenv("JWT_REFRESH_SECRET")

func GeraTokensESetaCookies(id int64, nome, nomeDeUsuario string, c echo.Context) error {
	accessToken, exp, err := GeraTokenAcesso(id, nome, nomeDeUsuario)

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)
	SetCookieUsuario(nomeDeUsuario, exp, c)

	refreshToken, exp, err := GeraTokenRefresh(id, nome, nomeDeUsuario)

	if err != nil {
		return err
	}

	SetCookieToken("refresh-token", refreshToken, exp, c)

	return nil
}

func GeraTokenAcesso(id int64, nome, nomeDeUsuario string) (string, time.Time, error) {
	expiraEm := time.Now().Add(1 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, expiraEm, []byte(ChaveDeAcesso))
}

func GeraTokenRefresh(id int64, nome, nomeDeUsuario string) (string, time.Time, error) {
	expiraEm := time.Now().Add(24 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, expiraEm, []byte(ChaveDeRefresh))
}

func GeraToken(id int64, nome, nomeDeUsuario string, expiraEm time.Time, chave []byte) (string, time.Time, error) {
	claims := &Claims{
		Id:            id,
		Nome:          nome,
		NomeDeUsuario: nomeDeUsuario,
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
		return fmt.Errorf("Token JWT inválido.")
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
	return c.Path() == "/v1/api/usuario/login" || (c.Path() == "/v1/api/usuario" && c.Request().Method == "POST") || strings.Contains(c.Path(), "/v1/api/docs")
}

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("usuario") == nil || PathWithNoAuthRequired(c) {
			return next(c)
		}

		nomeDeUsuario := c.Get("usuario").(*jwt.Token)

		claims := nomeDeUsuario.Claims.(*Claims)

		if time.Until(claims.RegisteredClaims.ExpiresAt.Time) < 15*time.Minute {
			refreshCookie, err := c.Cookie("refresh-token")

			if err == nil && refreshCookie != nil {
				token, err := jwt.ParseWithClaims(refreshCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(ChaveDeRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						return c.JSON(http.StatusUnauthorized, "")
					}
				}

				if token != nil && token.Valid {
					err = GeraTokensESetaCookies(nomeDeUsuario.Claims.(*Claims).Id, nomeDeUsuario.Claims.(*Claims).Nome, nomeDeUsuario.Claims.(*Claims).NomeDeUsuario, c)

					if err != nil {
						return c.JSON(http.StatusInternalServerError, "")
					}
				}
			}
		}

		return next(c)
	}
}
