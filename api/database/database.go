package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Conectar() (*sql.DB, error) {
	string := "root:Go7/flo2@/proje?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", string)
	if err != nil {
		log.Fatalf("Deu erro %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
