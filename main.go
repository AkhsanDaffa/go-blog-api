package main

import (
	"database/sql"
	"log"

	"go-blog-api/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.Println("Memulai aplikasi blog...")

	connStr := config.GetDBConnectionString()
	if connStr == "" {
		log.Fatal("Connection string tidak ditemukan")
	}

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke DB: %v", err)
	}

	log.Println("Berhasil terhubung ke Database")
}
