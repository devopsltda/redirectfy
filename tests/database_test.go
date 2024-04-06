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
			"1",
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
			"1",
		)
		assert.Error(t, err)
	})
	
	t.Run("Atualizar usuário do banco de dados", func(t *testing.T) {
		u.Create( //Cria usuario
			"53076281291",
			"Guilherme Bernardo",
			"GuilhermeBn",
			"bguilherme51@gmail.com",
			"Senha123",
			"01-01-2000",
			"1",
		)		
		err := u.Update(
			"09921080218",
			"Guilherme Lucas Pereira Bernardo",
			"GuilhermeBN",
			"bguilherem51@gmail.com",
			"senha-muito-complexa-aqui",
			"2000-10-31",
			"1",
		)
		assert.NoError(t, err)
	})
	
	t.Run("Atualizar usuário(COM OS MESMOS DADOS) do banco de dados", func(t *testing.T) {
		u.Create( //Cria usuario
			"09921080218",
			"Guilherme Lucas Pereira Bernardo",
			"GuilhermeBN",
			"bguilherem51@gmail.com",
			"senha-muito-complexa-aqui",
			"2000-10-31",
			"1",
		)		
		err := u.Update(
			"09921080218",
			"Guilherme Lucas Pereira Bernardo",
			"GuilhermeBN",
			"bguilherem51@gmail.com",
			"senha-muito-complexa-aqui",
			"2000-10-31",
			"1",
		)
		assert.Equal(t, err, nil)	
	})
	t.Run("Consultar 1 usuário do banco de dados", func(t *testing.T) {
		u.Create( //Cria usuario
			"53076281291",
			"Guilherme Bernardo",
			"GuilhermeBn",
			"bguilherme51@gmail.com",
			"Senha123",
			"01/01/2000",
			"1",
		)		
		_, err := u.ReadByNomeDeUsuario(
			"GuilhermeBn",
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
		u.Create( //Cria usuario
			"53076281291",
			"Guilherme Bernardo",
			"GuilhermeBn",
			"bguilherme51@gmail.com",
			"Senha123",
			"01/01/2000",
			"1",
		)		
		idUser, _, _, _, senhaUser, err := u.Login(
			"bguilherme51@gmail.com",
		)
		assert.NoError(t, err)
		t.Run("Troca Senha no banco de dados", func(t *testing.T) {
			err := u.TrocaSenha(
				idUser,
				senhaUser,
			)
			assert.NoError(t, err)
		})
	})
}

func TestDatabaseLinksRedirect(t *testing.T) {
	db := database.New()
	defer db.Close()
	l := models.LinkModel{DB: db}
	r := models.RedirecionadorModel{DB: db}
	t.Run("Criando redirecionador", func(t *testing.T) {
		_, err := r.Create(
			"redirectTest",
			"aa102930a22a",
			"1",
			1,
		)
		assert.NoError(t, err)
		t.Run("Criando Link no banco de dados", func(t *testing.T) {
			// Prepare the data for the Create method
			batchIdentifier := "batch1" // This should be a string that identifies the batch of links
			linksToInsert := []models.LinkToBatchInsert{
				{Nome: "Link1", Link: "https://example.com/link1", Plataforma: "Platform1"},
				{Nome: "Link2", Link: "https://example.com/link2", Plataforma: "Platform2"},
				// Add more LinkToBatchInsert objects as needed
			}
			
			// Call the Create method with the correct arguments
			err := l.Create(batchIdentifier, linksToInsert)
			assert.NoError(t, err)
		})
	})
}
