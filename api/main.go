package main

import (
	"projeto/database"
	"projeto/serve"
)

func main() {
	serve.Servidor()
	database.Conectar()
}
