package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"projeto/database"
	"projeto/products/respostas"
	"projeto/users/autenticacao"
	"projeto/users/modelos"
	"projeto/users/repositorios"
	"strconv"
	"strings"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := database.Conectar()
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

// BuscarUsuario busca um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := chi.URLParam(r, "id")

	usuarioID, erro := strconv.ParseUint(parametros, 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

// AtualizarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := chi.URLParam(r, "id")
	usuarioID, erro := strconv.ParseUint(parametros, 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := chi.URLParam(r, "id")
	usuarioID, erro := strconv.ParseUint(parametros, 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
