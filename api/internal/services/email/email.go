package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"text/template"

	_ "github.com/joho/godotenv/autoload"
)

type Email struct {
	addr string
	from string
	auth smtp.Auth
}

var uriToSendTo = os.Getenv("EMAIL_URI_TO_SEND_TO")

func New() *Email {
	username := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")
	port := os.Getenv("EMAIL_PORT")
	host := os.Getenv("EMAIL_HOST")
	from := os.Getenv("EMAIL_FROM")

	return &Email{
		addr: host + ":" + port,
		from: from,
		auth: smtp.PlainAuth("", username, password, host),
	}
}

func (e *Email) SendValidacao(id int64, nome, valor, email string) error {
	var messageTemplate strings.Builder
	defer messageTemplate.Reset()

	message := `From: Redirectfy <{{ .From }}>
To: {{ .Nome }} <{{ .Email }}>
Subject: Olá, {{ .Nome }}! Confirme seu endereço de email
Message-ID: <{{ .Id }}no-reply@redirectfy.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Confirmação de e-mail</title>
</head>
<body
    style="text-align: justify; height: 100vh; font-size: 20px; width: 100%; display: flex; justify-content: center; align-items: center; flex-direction: column; font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f4f4; padding: 10px;">
    <div
        style=" max-width: 600px; margin: 20px auto; padding: 20px; background-color: #fff; border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
        <div style="width:100%; display: flex; justify-content: center; gap: 10px;">
                <svg width="50" viewBox="0 0 34 32" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M3.60194 0H0.390869V3.21107V28.7889V32H3.60194L33.6089 32V28.7889H3.60194V3.21107H30.3978V23.1419H33.6089V0H33.6089H30.3978H3.60194ZM22.3148 10.9619H14.3425H11.1314V14.173V22.2561H14.3425V16.1411L21.6504 23.0032V23.1419H21.7981L21.8307 23.1726L21.8596 23.1419H30.3978V19.9308H23.0693L16.9374 14.173H22.3148V10.9619Z" fill="#35B5AE"/>
                </svg>
                <p style="color: #000;font-size: 40px; font-weight: bold; padding: 4px;">Redirectfy</p>
        </div>
        <h2 style="color: #000; font-size: 20px;">Confirme seu endereço de e-mail</h2>
        <p style="color: #666; ">Olá {{ .Nome }}!</p>
        <p style="color: #666;">Obrigado por se cadastrar no Redirectfy. Para garantir a segurança da sua conta,
            precisamos que você confirme seu endereço de e-mail. Por favor, clique no botão abaixo para verificar seu
            endereço de e-mail.</p>
        <div style="width: 100%; display: flex; justify-content: center;">
            <a href="{{ .URI }}finishSignup/{{ .Valor }}"
                style="font-size: 25px; padding: 10px 20px; background-color: #35B5AE; color: #fff; text-decoration: none; border-radius: 3px;">
                Confirmar e-mail
            </a>
        </div>
        <p style="color: #666;"">Caso não consiga pelo botão acima, utilize o link abaixo:</p>
        <p style="color: #666;"">{{ .URI }}finishSignup/{{ .Valor }}</p>
        <p style=" color: #666;">Se você não criou uma conta, por favor desconsidere este e-mail.</p>
        <p style="color: #666;">Atenciosamente,<br>Equipe do Redirectfy</p>
    </div>
    <div style="text-align: center; color: #999; font-size: 16px;">
        Não responda este e-mail.
    </div>
</body>
</html>`

	tmpl, err := template.New("").Parse(message)

	if err != nil {
		return err
	}

	err = tmpl.Execute(&messageTemplate, map[string]string{
		"Id":    fmt.Sprint(id),
		"Nome":  nome,
		"Valor": valor,
		"Email": email,
		"From":  e.from,
		"URI":   uriToSendTo,
	})

	if err != nil {
		return err
	}

	return smtp.SendMail(e.addr, e.auth, e.from, []string{email}, []byte(messageTemplate.String()))
}

func (e *Email) SendTrocaDeSenha(id int64, nome, valor, email string) error {
	var messageTemplate strings.Builder
	defer messageTemplate.Reset()

	message := `From: Redirectfy <{{ .From }}>
To: {{ .Nome }} <{{ .Email }}>
Subject: Olá, {{ .Nome }}! Confirme seu endereço de email
Message-ID: <{{ .Id }}no-reply@redirectfy.com>
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Redefinição de senha</title>
</head>
<body
    style="text-align: justify; height: 100vh; font-size: 20px; width: 100%; display: flex; justify-content: center; align-items: center; flex-direction: column; font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f4f4f4; padding: 10px;">
    <div
        style=" max-width: 600px; margin: 20px auto; padding: 20px; background-color: #fff; border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
        <div style="width:100%; display: flex; justify-content: center; gap: 10px;">
                <svg width="50" viewBox="0 0 34 32" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M3.60194 0H0.390869V3.21107V28.7889V32H3.60194L33.6089 32V28.7889H3.60194V3.21107H30.3978V23.1419H33.6089V0H33.6089H30.3978H3.60194ZM22.3148 10.9619H14.3425H11.1314V14.173V22.2561H14.3425V16.1411L21.6504 23.0032V23.1419H21.7981L21.8307 23.1726L21.8596 23.1419H30.3978V19.9308H23.0693L16.9374 14.173H22.3148V10.9619Z" fill="#35B5AE"/>
                </svg>
                <p style="color: #000;font-size: 40px; font-weight: bold; padding: 4px;">Redirectfy</p>
        </div>
        <h2 style="color: #000; font-size: 20px;">Redefinição de senha</h2>
        <p style="color: #666; ">Olá {{ .Nome }}!</p>
        <p style="color: #666;">Recebemos uma solicitação de redefinição de senha para a sua conta no Redirectfy. Para fazer isso, basta clicar no botão abaixo para redefinir sua senha.</p>
        <div style="width: 100%; display: flex; justify-content: center;">
            <a href="{{ .URI }}finishSignup/{{ .Valor }}"
                style="font-size: 25px; padding: 10px 20px; background-color: #35B5AE; color: #fff; text-decoration: none; border-radius: 3px;">
                Redefinir senha
            </a>
        </div>
        <p style="color: #666;"">Caso não consiga pelo botão acima, utilize o link abaixo:</p>
        <p style="color: #666;"">{{ .URI }}finishSignup/{{ .Valor }}</p>
        <p style=" color: #666;">Se não foi você, basta ignorar esse e-mail.</p>
        <p style="color: #666;">Atenciosamente,<br>Equipe do Redirectfy</p>
    </div>
    <div style="text-align: center; color: #999; font-size: 16px;">
        Não responda este e-mail.
    </div>
</body>
</html>`

	tmpl, err := template.New("").Parse(message)

	if err != nil {
		return err
	}

	err = tmpl.Execute(&messageTemplate, map[string]string{
		"Id":    fmt.Sprint(id),
		"Nome":  nome,
		"Valor": valor,
		"Email": email,
		"From":  e.from,
		"URI":   uriToSendTo,
	})

	if err != nil {
		return err
	}

	return smtp.SendMail(e.addr, e.auth, e.from, []string{email}, []byte(messageTemplate.String()))
}
