package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projeto/database"
	"projeto/products/respostas"
	"projeto/users/autenticacao"
	"projeto/users/modelos"
	"projeto/users/repositorios"
	"projeto/users/seguranca"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	requisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuarios modelos.Usuario
	if err = json.Unmarshal(requisicao, &usuarios); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorios.NovoRepositorioDeUsuarios(db)
	usuariosSalvoNoBanco, err := repository.BuscarPorEmail(usuarios.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = seguranca.VerificarSenha(usuariosSalvoNoBanco.Senha, usuarios.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuariosSalvoNoBanco.ID)
	if err != nil {
		fmt.Sprintf("%s", err)
	}

	w.Write([]byte(fmt.Sprintf("LOGADO COM SUCESSO; EMAIL: %s, \n TOKEN - %s", usuarios.Email, token)))
}
