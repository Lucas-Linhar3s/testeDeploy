package modelos

import (
	"errors"
	"projeto/users/seguranca"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID           uint64    `json:"id,omitempty"`
	Nome         string    `json:"nome,omitempty"`
	Email        string    `json:"email,omitempty"`
	Senha        string    `json:"senha,omitempty"`
	ProdutoID    int       `json:"produtoID"`
	CriadoEm     time.Time `json:"CriadoEm,omitempty"`
	AtualizadoEm time.Time `json:"AtualizadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
