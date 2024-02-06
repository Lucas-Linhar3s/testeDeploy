package repositorio

import (
	"database/sql"
	"log"
	"projeto/products/modelos"
)

type Produtos struct {
	db *sql.DB
}

func NovoRepositorio(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// LISTAR TODOS OS PRODUTOS DO BANCO
func (u Produtos) Listar() ([]modelos.Produtos, error) {
	linhas, err := u.db.Query("SELECT * FROM produtos")
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	defer linhas.Close()

	var produtos []modelos.Produtos

	for linhas.Next() {
		var produto modelos.Produtos

		if err = linhas.Scan(&produto.ID, &produto.Nome_Produto, &produto.Valor, &produto.Quantidade, &produto.Dt_Cadastro, &produto.Dt_Atualizacao); err != nil {
			log.Printf("Deu Erro %v", err)
			return nil, err
		}

		produtos = append(produtos, produto)
	}
	return produtos, nil
}

// CRIAR NOVOS PRODUTOS
func (p Produtos) Criar(produtos modelos.Produtos) (uint64, error) {
	statement, err := p.db.Prepare("INSERT INTO produtos (produto, valor, quantidade) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	defer statement.Close()

	result, err := statement.Exec(produtos.Nome_Produto, produtos.Valor, produtos.Quantidade)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	return uint64(lastId), nil
}

// FILTRAR PRODUTOS POR ID
func (p Produtos) BuscaID(ID int64) (modelos.Produtos, error) {
	linhas, err := p.db.Query("SELECT * FROM produtos WHERE id = ?", ID)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	var produtos modelos.Produtos
	if linhas.Next() {
		if err = linhas.Scan(&produtos.ID, &produtos.Nome_Produto, &produtos.Valor, &produtos.Quantidade, &produtos.Dt_Cadastro, &produtos.Dt_Atualizacao); err != nil {
			log.Printf("Deu Erro %v", err)
			return modelos.Produtos{}, err
		}
	}
	return produtos, nil
}

// FILTRAR PRODUTOS POR PRECO
func (p Produtos) Filtrar(preco int64) ([]modelos.Produtos, error) {
	linhas, err := p.db.Query("SELECT * FROM produtos WHERE valor >= ?", preco)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	defer linhas.Close()

	var filtros []modelos.Produtos

	for linhas.Next() {
		var filtro modelos.Produtos
		if err = linhas.Scan(&filtro.ID, &filtro.Nome_Produto, &filtro.Valor, &filtro.Quantidade, &filtro.Dt_Cadastro, &filtro.Dt_Atualizacao); err != nil {
			log.Printf("Deu Erro %v", err)
			return nil, err
		}
		filtros = append(filtros, filtro)
	}
	return filtros, nil
}

func (p Produtos) Atualizar(ID int64, Produto modelos.Produtos) error {
	statement, erro := p.db.Prepare(
		"update produtos set produto = ?, valor = ?, quantidade = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Produto.Nome_Produto, Produto.Valor, Produto.Quantidade, ID); erro != nil {
		return erro
	}

	return nil
}

func (p Produtos) Deletar(ID int64) error {
	statement, err := p.db.Prepare("DELETE FROM produtos WHERE id = ?")
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		log.Printf("Deu Erro %v", err)
	}
	return nil
}
