package database

import (
	"fmt"
	"log"
	
	"diet_diary/internal/config"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(cfg *config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnString())
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully!")

	return db, nil
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed.")
	}
}
