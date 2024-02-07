package serve

import (
	"fmt"
	"log"
	"net/http"
	"projeto/middlewares"
	"projeto/products/controller"
	"projeto/users/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func Servidor() {
	r := chi.NewRouter()

	// IMPRIME LOGS NO CONSOLE SOBRE REQUISIÇÕES DE ROTAS
	r.Use(middleware.Logger)

	// HABILITA O CORS EM TODOS OS NAVEGADORES
	r.Use(cors.AllowAll().Handler)
	// ROTAS DE PRODUTOS
	r.Group(func(r chi.Router) {
		r.Get("/subscribe", controller.List)
		r.Post("/teste", controller.TesteTMS)                                          // LISTAR TODOS OS PRODUTOS DO BANCO
		r.Get("/produtos/listar/{id}", middlewares.Autenticar(controller.BuscarID))    // FILTRAR PRODUTOS POR ID
		r.Get("/produtos/filtrar/{preco}", middlewares.Autenticar(controller.Filtrar)) // FILTRAR PRODUTOS POR PREÇO
		r.Post("/produtos/criar", controller.Criar)                                    // CRIA NOVOS PRODUTOS
		r.Put("/produtos/atualizar/{id}", controller.Atualizar)                        // ATUALIZAR PRODUTOS
		r.Delete("/produtos/deletar/{id}", controller.Deletar)                         // DELETAR PRODUTOS
	})

	// ROTAS DE USUÁRIOS PRIVADA
	r.Group(func(r chi.Router) {
		r.Get("/usuarios/listar", middlewares.Autenticar(controllers.BuscarUsuarios))           // LISTA TODOS OS USUÁRIOS
		r.Get("/usuarios/listar/{id}", middlewares.Autenticar(controllers.BuscarUsuario))       // FILTRAR USUÁRIOS POR ID
		r.Delete("/usuarios/deletar/{id}", middlewares.Autenticar(controllers.DeletarUsuario))  // DELETAR USUÁRIOS
		r.Put("/usuarios/atualizar/{id}", middlewares.Autenticar(controllers.AtualizarUsuario)) // ATUALIZAR USUÁRIOS
	})

	// ROTAS DE USUÁRIOS PUBLICA
	r.Group(func(r chi.Router) {
		r.Post("/usuarios/criar", controllers.CriarUsuario) // CRIA NOVOS USUÁRIOS
		r.Post("/usuarios/login", controllers.Login)        // ROTA DE LOGIN
	})

	fmt.Println("SERVIDOR RODANDO NA PORTA 3333")
	log.Fatal(http.ListenAndServe(":3333", r))
}
