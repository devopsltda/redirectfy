package email

import (
	"net/smtp"
	"os"

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

func SendEmailValidacao(to []string) error {
	message := []byte(`From: Redirectify <no-reply@redirectify.com>
To: Eduardo Machado <edu.hen.fm@gmail.com>
Subject: Olá, Eduardo Machado! Confirme seu endereço de email
Date: Tue, 28 Jul 2015 20:02:14 +0000
Message-ID: <12345no-reply@redirectify.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline
Received: ;

<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
  </head>
  <div>
    <h1>Confirme seu email</h1>
    <p>Clique <a href="https://google.com">aqui</a>.</p>
    <p>
    Caso não consiga, acesse diretamente pelo link 
    <a href="https://google.com">https://google.com</a>
    .</p>
    <strong>Não responda este email.</strong>
  </div>
<html>
`)

	return smtp.SendMail(host+":"+port, auth, from, to, message)
}

func SendEmailTrocaDeSenha(to []string) error {
	message := []byte(`From: Redirectify <no-reply@redirectify.com>
To: Eduardo Machado <edu.hen.fm@gmail.com>
Subject: Olá, Eduardo Machado! Confirme seu endereço de email
Date: Tue, 28 Jul 2015 20:02:14 +0000
Message-ID: <12345no-reply@redirectify.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline
Received: ;

<!DOCTYPE html>
<html>
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
  </head>
  <div>
    <h1>Confirme seu email</h1>
    <p>Clique <a href="https://google.com">aqui</a>.</p>
    <p>
    Caso não consiga, acesse diretamente pelo link 
    <a href="https://google.com">https://google.com</a>
    .</p>
    <strong>Não responda este email.</strong>
  </div>
<html>
`)

	return smtp.SendMail(host+":"+port, auth, from, to, message)
}
