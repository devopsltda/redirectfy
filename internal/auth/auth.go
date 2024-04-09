package auth

import (
	"fmt"
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
	return (c.Request().URL.Path == "/usuarios/login" && c.Request().Method == "POST") ||
		(c.Request().URL.Path == "/usuarios" && c.Request().Method == "POST") ||
		(c.Path() == "/usuarios/troca_de_senha/:valor" && c.Request().Method == "PATCH") ||
		(c.Path() == "/usuarios/:nome_de_usuario/troca_de_senha" && c.Request().Method == "POST") ||
		(c.Path() == "/docs/*" && c.Request().Method == "GET") ||
		(c.Path() == "/autenticacao/:valor" && c.Request().Method == "PATCH") ||
		(c.Path() == "/planos_de_assinatura" && c.Request().Method == "GET") ||
		(c.Path() == "/planos_de_assinatura/:nome" && c.Request().Method == "GET") ||
		(c.Path() == "/usuarios_temporarios" && c.Request().Method == "POST") ||
		(c.Path() == "/usuarios/criar_permanente/:valor" && c.Request().Method == "POST") ||
		(c.Path() == "/redirecionadores/:codigo_hash/rehash" && c.Get("usuario") != nil && c.Get("usuario").(*jwt.Token).Claims.(*Claims).PlanoDeAssinatura == "Super Pro")
}

func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if PathWithNoAuthRequired(c) {
			return next(c)
		}

		if c.Get("usuario") == nil {
			return utils.ErroAssinaturaJWT
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
						return utils.ErroAssinaturaJWT
					}
				}

				if token != nil && token.Valid {
					err = GeraTokensESetaCookies(claims.Id, claims.Nome, claims.NomeDeUsuario, claims.PlanoDeAssinatura, c)

					if err != nil {
						return utils.ErroAssinaturaJWT
					}
				}
			}
		}

		return next(c)
	}
}
