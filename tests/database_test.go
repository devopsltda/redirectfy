package tests

import (
	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseUsuario(t *testing.T) {
	db := database.New()
	defer db.Close()
	u := models.UsuarioModel{DB: db}

	t.Run("Adicionar usuário ao banco de dados", func(t *testing.T) {
		_, err := u.Create(
			"02958985261",
			"Eduardo Henrique Freire Machado",
			"edu_hen_fm",
			"edu.hen.fm@gmail.com",
			"senha-muito-complexa-aqui",
			"2001-08-30",
			1,
		)
		assert.NoError(t, err)
	})

	t.Run("Adicionar usuário DUPLICADO(msm cpf) ao banco de dados", func(t *testing.T) {
		_, err := u.Create(
			"02958985261",
			"Eduardo Henrique Freire Machado",
			"edu_hen_fm",
			"edu.hen.fm@gmail.com",
			"senha-muito-complexa-aqui",
			"2001-08-30",
			1,
		)
		assert.NoError(t, err)
	})

	t.Run("Atualizar usuário do banco de dados", func(t *testing.T) {
		err := u.Update(
			"09921080218",
			"Guilherme Lucas Pereira Bernardo",
			"GuilhermeBN",
			"bguilherem51@gmail.com",
			"senha-muito-complexa-aqui",
			"2000-10-31",
			1,
		)
		assert.NoError(t, err)
	})

	t.Run("Atualizar usuário(COM OS MESMOS DADOS) do banco de dados", func(t *testing.T) {
		err := u.Update(
			"09921080218",
			"Guilherme Lucas Pereira Bernardo",
			"GuilhermeBN",
			"bguilherem51@gmail.com",
			"senha-muito-complexa-aqui",
			"2000-10-31",
			1,
		)
		assert.NoError(t, err)
	})

	t.Run("Consultar 1 usuário do banco de dados", func(t *testing.T) {
		_, err := u.ReadByNomeDeUsuario(
			"GuilhermeBN",
		)
		assert.NoError(t, err)
	})

	t.Run("Consultar usuários do banco de dados", func(t *testing.T) {
		_, err := u.ReadAll()
		assert.NoError(t, err)
	})

	t.Run("Deletar Usuario do banco de dados", func(t *testing.T) {
		err := u.Remove(
			"GuilhermeBN",
		)
		assert.NoError(t, err)
	})
}

func TestDatabaseUsuarioAuth(t *testing.T) {
	db := database.New()
	defer db.Close()
	u := models.UsuarioModel{DB: db}

	t.Run("Login do usuario no Banco de dados", func(t *testing.T) {
		_, _, _, _, _, err := u.Login(
			"bguilherme51@gmail.com",
			"GuilhermeBN",
		)
		assert.NoError(t, err)
	})

	t.Run("Validação de email no banco de dados", func(t *testing.T) {
		err := u.Autenticado(
			1,
		)
		assert.NoError(t, err)
	})
}

func TestDatabaseLinksRedirect(t *testing.T) {
	db := database.New()
	defer db.Close()
	l := models.LinkModel{DB: db}
	t.Run("Criando redirecionador", func(t *testing.T) {

	})
	t.Run("Criando Link no banco de dados", func(t *testing.T) {
		err := l.Create(
			"rock garai",
			"saodaido",
			"askjdasjk",
			"askdkjasd",
		)
		assert.NoError(t, err)
	})

	t.Run("Pegar todas os links", func(t *testing.T) {
		_, err := l.ReadAll(
			"osdsaoid10i912id0asd",
		)
		assert.NoError(t, err)
	})
}
