package modelos

import "time"

type Produtos struct {
	ID             int64     `json:"id,string"`
	Nome_Produto   string    `json:"nome_produto"`
	Valor          float32   `json:"valor,string"`
	Quantidade     int64     `json:"quantidade,string"`
	Dt_Cadastro    time.Time `json:"dt_Cadastro,omitempty"`
	Dt_Atualizacao time.Time `json:"dt_Atualizacao,omitempty"`
}
