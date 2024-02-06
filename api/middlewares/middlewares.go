package middlewares

import (
	"fmt"
	"net/http"
	"projeto/products/respostas"
	"projeto/users/autenticacao"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			fmt.Println("MIDDLEWARE!")
			return
		}
		proximaFuncao(w, r)
	}
}
