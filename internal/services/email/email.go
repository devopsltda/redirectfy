package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"text/template"

	_ "github.com/joho/godotenv/autoload"
)

var (
	username = os.Getenv("EMAIL_USERNAME")
	password = os.Getenv("EMAIL_PASSWORD")
	port     = os.Getenv("EMAIL_PORT")
	host     = os.Getenv("EMAIL_HOST")
	from     = os.Getenv("EMAIL_FROM")
	auth     = smtp.CRAMMD5Auth(username, password)
)

func SendEmailValidacao(id int64, nome, valor, email string) error {
	var messageTemplate strings.Builder
	defer messageTemplate.Reset()

	message := `From: Redirectify <{{ .From }}>
To: {{ .Nome }} <{{ .Email }}>
Subject: Olá, {{ .Nome }}! Confirme seu endereço de email
Message-ID: <{{ .Id }}no-reply@redirectify.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
  </head>
  <div>
    <h1>Confirme seu email</h1>
    <p>Clique <a href="https://localhost:8080/v1/api/autenticacao/{{ .Valor }}">aqui</a>.</p>
    <p>
    Caso não consiga, acesse diretamente pelo link 
		<a href="https://localhost:8080/v1/api/autenticacao/{{ .Valor }}">https://localhost:8080/v1/api/autenticacao/{{ .Valor }}</a>
    .</p>
    <strong>Não responda este email.</strong>
  </div>
<html>`

	tmpl, err := template.New("").Parse(message)

	if err != nil {
		return err
	}

	err = tmpl.Execute(&messageTemplate, map[string]string{
		"Id":    fmt.Sprint(id),
		"Nome":  nome,
		"Valor": valor,
		"Email": email,
		"From":  from,
	})

	if err != nil {
		return err
	}

	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(messageTemplate.String()))
}

func SendEmailTrocaDeSenha(id int64, nome, valor, email string) error {
	var messageTemplate strings.Builder
	defer messageTemplate.Reset()

	message := `From: Redirectify <{{ .From }}>
To: {{ .Nome }} <{{ .Email }}>
Subject: Olá, {{ .Nome }}! Confirme seu endereço de email
Message-ID: <{{ .Id }}no-reply@redirectify.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
  </head>
  <div>
    <h1>Confirme seu email</h1>
    <p>Clique <a href="https://localhost:8080/v1/api/usuarios/troca_de_senha/{{ .Valor }}">aqui</a>.</p>
    <p>
    Caso não consiga, acesse diretamente pelo link 
		<a href="https://localhost:8080/v1/api/usuarios/troca_de_senha/{{ .Valor }}">https://localhost:8080/v1/api/usuarios/troca_de_senha/{{ .Valor }}</a>
    .</p>
    <strong>Não responda este email.</strong>
  </div>
<html>`

	tmpl, err := template.New("").Parse(message)

	if err != nil {
		return err
	}

	err = tmpl.Execute(&messageTemplate, map[string]string{
		"Id":    fmt.Sprint(id),
		"Nome":  nome,
		"Valor": valor,
		"Email": email,
		"From":  from,
	})

	if err != nil {
		return err
	}

	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(messageTemplate.String()))
}
