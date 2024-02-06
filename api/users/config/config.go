package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
