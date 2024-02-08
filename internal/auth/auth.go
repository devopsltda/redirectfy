package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	Nome string `json:"nome"`
	jwt.RegisteredClaims
}

var ChaveDeAcesso = os.Getenv("JWT_SECRET")
var ChaveDeRefresh = os.Getenv("JWT_REFRESH_SECRET")

func GeraTokensESetaCookies(nomeDeUsuario string, c echo.Context) error {
	accessToken, exp, err := GeraTokenAcesso(nomeDeUsuario)

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)
	SetCookieUsuario(nomeDeUsuario, exp, c)

	refreshToken, exp, err := GeraTokenRefresh(nomeDeUsuario)

	if err != nil {
		return err
	}

	SetCookieToken("refresh-token", refreshToken, exp, c)

	return nil
}

func GeraTokenAcesso(nomeDeUsuario string) (string, time.Time, error) {
	expiraEm := time.Now().Add(1 * time.Hour)

	return GeraToken(nomeDeUsuario, expiraEm, []byte(ChaveDeAcesso))
}

func GeraTokenRefresh(nomeDeUsuario string) (string, time.Time, error) {
	expiraEm := time.Now().Add(24 * time.Hour)

	return GeraToken(nomeDeUsuario, expiraEm, []byte(ChaveDeRefresh))
}

func GeraToken(nomeDeUsuario string, expiraEm time.Time, chave []byte) (string, time.Time, error) {
	claims := &Claims{
		Nome: nomeDeUsuario,
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
		return fmt.Errorf("Invalid JWT Token")
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

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("usuario") == nil {
			return next(c)
		}

		nomeDeUsuario := c.Get("usuario").(*jwt.Token)

		claims := nomeDeUsuario.Claims.(*Claims)

		if time.Unix(claims.RegisteredClaims.ExpiresAt.Unix(), 0).Sub(time.Now()) < 15*time.Minute {
			refreshCookie, err := c.Cookie("refresh-token")

			if err == nil && refreshCookie != nil {
				token, err := jwt.ParseWithClaims(refreshCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(ChaveDeRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.JSON(http.StatusUnauthorized, "")
					}
				}

				if token != nil && token.Valid {
					err = GeraTokensESetaCookies(nomeDeUsuario.Claims.(*Claims).Nome, c)

					if err != nil {
						return c.JSON(http.StatusInternalServerError, "")
					}
				}
			}
		}

		return next(c)
	}
}
