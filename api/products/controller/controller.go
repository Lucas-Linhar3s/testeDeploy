package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projeto/database"
	modelos2 "projeto/products/modelos"
	repositorio3 "projeto/products/repositorio"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Detail  string `json:"detail"`
	Message string `json:"message"`
}

type Response struct {
	DeviceID *string `json:"deviceId"`
	Action   *string `json:"action"`
	Value    *string `json:"value"`
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	errorResponse := ErrorResponse{
		Code:    3304,
		Detail:  "Action refused",
		Message: "Request refused due to a technical error. Please try again later.",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorResponse)
}

func TesteTMS(w http.ResponseWriter, r *http.Request) {

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	var response []Response

	if err = json.Unmarshal(request, &response); err != nil {
		log.Printf("Deu Erro %v", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Criar(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	var produtos modelos2.Produtos

	if err = json.Unmarshal(request, &produtos); err != nil {
		log.Printf("Deu Erro %v", err)
	}
	db, err := database.Conectar()
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	repositorio := repositorio3.NovoRepositorio(db)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	produto, err := repositorio.Criar(produtos)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("PRODUTO CRIADO COM SUCESSO: PRODUTO: %s, ID: %v", produtos.Nome_Produto, produto)))
}

func BuscarID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := chi.URLParam(r, "id")

	ProdutoID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	db, err := database.Conectar()
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	repositorio := repositorio3.NovoRepositorio(db)
	produtos, err := repositorio.BuscaID(ProdutoID)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produtos)
}

func Filtrar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "preco")

	Preco, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	repositorio := repositorio3.NovoRepositorio(db)
	produto, err := repositorio.Filtrar(Preco)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	json.NewEncoder(w).Encode(produto)
	w.WriteHeader(http.StatusOK)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	parametros := chi.URLParam(r, "id")
	produID, erro := strconv.ParseInt(parametros, 10, 64)
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	var produto modelos2.Produtos
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	repositorio := repositorio3.NovoRepositorio(db)
	if erro = repositorio.Atualizar(produID, produto); erro != nil {
		log.Printf("Deu Erro %v", erro)
		return
	}

	w.Write([]byte(fmt.Sprintf("PRODUTO: %s, ATUALIZADO", produto.Nome_Produto)))
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")

	ProduID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}

	db, err := database.Conectar()
	if err != nil {
		log.Printf("Deu Erro %v", err)
	}
	repositorio := repositorio3.NovoRepositorio(db)
	if err = repositorio.Deletar(ProduID); err != nil {
		log.Printf("Deu Erro %v", err)
	}

	w.Write([]byte(fmt.Sprintf("PRODUTO COM ID: %v, FOI EXCLUIDO", ProduID)))
}
