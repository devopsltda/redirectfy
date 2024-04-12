package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"redirectfy/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ChaveDeAcesso  = os.Getenv("JWT_SECRET")
	ChaveDeRefresh = os.Getenv("JWT_REFRESH_SECRET")
)

type Claims struct {
	Id                int64  `json:"id"`
	Nome              string `json:"nome"`
	NomeDeUsuario     string `json:"nome_de_usuario"`
	PlanoDeAssinatura string `json:"plano_de_assinatura"`
	jwt.RegisteredClaims
}

func GeraTokensESetaCookies(id int64, nome, nomeDeUsuario, planoDeAssinatura string, c echo.Context) error {
	accessToken, exp, err := GeraTokenAcesso(id, nome, nomeDeUsuario, planoDeAssinatura)

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)
	SetCookieUsuario(nomeDeUsuario, exp, c)

	refreshToken, exp, err := GeraTokenRefresh(id, nome, nomeDeUsuario, planoDeAssinatura)

	if err != nil {
		return err
	}

	SetCookieToken("refresh-token", refreshToken, exp, c)

	return nil
}

func GeraTokenAcesso(id int64, nome, nomeDeUsuario, planoDeAssinatura string) (string, time.Time, error) {
	expiraEm := time.Now().Add(1 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, planoDeAssinatura, expiraEm, []byte(ChaveDeAcesso))
}

func GeraTokenRefresh(id int64, nome, nomeDeUsuario, planoDeAssinatura string) (string, time.Time, error) {
	expiraEm := time.Now().Add(24 * time.Hour)

	return GeraToken(id, nome, nomeDeUsuario, planoDeAssinatura, expiraEm, []byte(ChaveDeRefresh))
}

func GeraToken(id int64, nome, nomeDeUsuario, planoDeAssinatura string, expiraEm time.Time, chave []byte) (string, time.Time, error) {
	claims := &Claims{
		Id:                id,
		Nome:              nome,
		NomeDeUsuario:     nomeDeUsuario,
		PlanoDeAssinatura: planoDeAssinatura,
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
		return errors.New("O seu token de autenticação é inválido.")
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
	return (c.Request().URL.Path == "/u/login" && c.Request().Method == "POST") ||
		(c.Path() == "/u/change_password/:hash" && c.Request().Method == "PATCH") ||
		(c.Path() == "/u/:username/change_password" && c.Request().Method == "POST") ||
		(c.Path() == "/docs/*" && c.Request().Method == "GET") ||
		(c.Path() == "/pricing" && c.Request().Method == "GET") ||
		(c.Path() == "/pricing/:name" && c.Request().Method == "GET") ||
		(c.Path() == "/kirvano" && c.Request().Method == "POST") ||
		(c.Path() == "/kirvano/to_user/:hash" && c.Request().Method == "POST") ||
		(c.Path() == "/r/:hash" && c.Request().Method == "GET") || 
		(c.Path() == "/r/:hash/refresh" && c.Get("usuario") != nil && c.Get("usuario").(*jwt.Token).Claims.(*Claims).PlanoDeAssinatura == "Super Pro")
}

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if PathWithNoAuthRequired(c) {
			return next(c)
		}

		if c.Get("usuario") == nil {
			return utils.Erro(http.StatusBadRequest, "Você não contém um ou mais dos cookies necessários para autenticação.")
		}

		nomeDeUsuario := c.Get("usuario").(*jwt.Token)
		nomeDeUsuarioCookie, err := c.Cookie("usuario")

		if err != nil {
			return utils.Erro(http.StatusBadRequest, "Você não contém um ou mais dos cookies necessários para autenticação.")
		}

		claims := nomeDeUsuario.Claims.(*Claims)

		if claims.NomeDeUsuario != nomeDeUsuarioCookie.Value {
			return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação é inválido.")
		}

		if time.Until(claims.RegisteredClaims.ExpiresAt.Time) < 15*time.Minute {
			refreshCookie, err := c.Cookie("refresh-token")

			if err == nil && refreshCookie != nil {
				token, err := jwt.ParseWithClaims(refreshCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(ChaveDeRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação é inválido.")
					}
				}

				if token != nil && token.Valid {
					err = GeraTokensESetaCookies(claims.Id, claims.Nome, claims.NomeDeUsuario, claims.PlanoDeAssinatura, c)

					if err != nil {
						return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação é inválido.")
					}
				}
			}
		}

		return next(c)
	}
}
