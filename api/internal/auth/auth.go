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

	limiteAccess = 1 * time.Hour
	limiteRefresh = 2 * time.Hour
	limiteParaCriacaoDeNovoAccess = 15 * time.Minute
)

// Claims representa as informações que são armazenadas nos tokens JWT.
type Claims struct {
	Id                int64  `json:"id"`
	Nome              string `json:"nome"`
	NomeDeUsuario     string `json:"nome_de_usuario"`
	PlanoDeAssinatura string `json:"plano_de_assinatura"`
	jwt.RegisteredClaims
}

// GeraTokensESetaCookies gera os tokens 'access' e 'refresh', os transforma em
// cookies e os insere no contexto fornecido.
func GeraTokensESetaCookies(id int64, nome, nomeDeUsuario, planoDeAssinatura string, c echo.Context) error {
	accessToken, exp, err := GeraToken(id, nome, nomeDeUsuario, planoDeAssinatura, time.Now().Add(limiteAccess), []byte(ChaveDeAcesso))

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)

	refreshToken, exp, err := GeraToken(id, nome, nomeDeUsuario, planoDeAssinatura, time.Now().Add(limiteRefresh), []byte(ChaveDeRefresh))

	if err != nil {
		return err
	}

	SetCookieToken("refresh-token", refreshToken, exp, c)

	return nil
}

// GeraTokensESetaCookiesSemRefresh gera o token 'access', o transforma em
// cookies e o insere no contexto fornecido.
func GeraTokensESetaCookiesSemRefresh(id int64, nome, nomeDeUsuario, planoDeAssinatura string, c echo.Context) error {
	accessToken, exp, err := GeraToken(id, nome, nomeDeUsuario, planoDeAssinatura, time.Now().Add(limiteAccess), []byte(ChaveDeAcesso))

	if err != nil {
		return err
	}

	SetCookieToken("access-token", accessToken, exp, c)

	return nil
}

// GeraToken gera um token JWT contendo o id, nome, nome de usuário, plano de
// assinatura, data de expiração e chave. Ele retorna o token, o tempo de
// expiração (pode ser usado para adicionar o tempo de expiração em cookies,
// por exemplo) e um erro de assinatura, caso não tenha sido possível assinar
// o token.
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

// VerificaToken verifica se o token fornecido é válido.
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

// SetCookieToken adiciona um token ao contexto com o nome, token JWT e data de
// expiração fornecidos.
func SetCookieToken(nome, token string, expiraEm time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = nome
	cookie.Value = token
	cookie.Expires = expiraEm
	cookie.Path = "/"

	c.SetCookie(cookie)
}


// PathWithNoAuthRequired verifica se o caminho atual do contexto exige ou não
// autorização (ou seja, se é necessário ou não conferir o token de acesso do
// contexto).
func PathWithNoAuthRequired(c echo.Context) bool {
	return (c.Request().URL.Path == "/api/u/login" && c.Request().Method == "POST") ||
		(c.Path() == "/api/u/change_password/:hash" && c.Request().Method == "PATCH") ||
		(c.Path() == "/api/u/:username/change_password" && c.Request().Method == "POST") ||
		(c.Path() == "/api/docs/*" && c.Request().Method == "GET") ||
		(c.Path() == "/api/pricing" && c.Request().Method == "GET") ||
		(c.Path() == "/api/pricing/:name" && c.Request().Method == "GET") ||
		(c.Path() == "/api/kirvano" && c.Request().Method == "POST") ||
		(c.Path() == "/api/kirvano/to_user/:hash" && c.Request().Method == "POST") ||
		(c.Path() == "/api/to/:hash" && c.Request().Method == "GET")
}

// TokenRefreshMiddleware verifica se o token de acesso do contexto está dentro
// do limite de tempo 'limiteParaCriacaoDeNovoAccess' para gerar um novo token
// de acesso sem que o usuário tenha que fazer login novamente. Caso o token de
// refresh seja inválido, um erro é devolvido.
func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if PathWithNoAuthRequired(c) || c.Get("usuario") == nil {
			return next(c)
		}

		// Note que o token, ao chegar até aqui, já passou pelo middleware de
		// autenticação, logo ele existe e não está expirado.
		token := c.Get("usuario").(*jwt.Token)
		claims := token.Claims.(*Claims)
		
		timeUntilAccessExpiration := time.Until(claims.RegisteredClaims.ExpiresAt.Time)

		if timeUntilAccessExpiration < limiteParaCriacaoDeNovoAccess {
			refreshCookie, err := c.Cookie("refresh-token")

			if err == nil && refreshCookie != nil {
				token, err := jwt.ParseWithClaims(refreshCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(ChaveDeRefresh), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						utils.DebugLog("TokenRefreshMiddleware", "Erro na assinatura do token JWT", err)
						return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação é inválido.")
					}
				}

				if token != nil && token.Valid {
					// É importante que o refresh não seja setado novamente para que este
					// possa expirar no 'limiteRefresh' estabelecido quando o token foi
					// criado.
					err = GeraTokensESetaCookiesSemRefresh(claims.Id, claims.Nome, claims.NomeDeUsuario, claims.PlanoDeAssinatura, c)

					if err != nil {
						utils.DebugLog("TokenRefreshMiddleware", "Erro na validação do token de autenticação", err)
						return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação é inválido.")
					}
				}
			}
		}

		return next(c)
	}
}
