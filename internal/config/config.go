package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Srv *Server
	DB  *DB
}

type Server struct {
	host string
	port int
}

func (srv *Server) Host() string {
	return fmt.Sprintf("%s:%d", srv.host, srv.port)
}

type DB struct {
	source   string
	name     string
	user     string
	password string
	address  string
	port     int
}

func (cfg *DB) ConnString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", cfg.source, cfg.user, cfg.password, cfg.address, cfg.port, cfg.name)
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cfg := &Config{
		Srv: &Server{
			host: getEnvAsString("WEBSERVER_HOST", "localhost"),
			port: getEnvAsInteger("WEBSERVER_PORT", 8080),
		},
		DB: &DB{
			source:   getEnvAsString("DB_SOURCE", "postgres"),
			name:     getEnvAsString("DB_NAME", "db"),
			user:     getEnvAsString("DB_USER", "admin"),
			password: getEnvAsString("DB_PASSWORD", ""),
			address:  getEnvAsString("DB_ADDRESS", "localhost"),
			port:     getEnvAsInteger("DB_PORT", 5432),
		},
	}
	return cfg, nil
}

func getEnvAsString(key, defaultVal string) string {
	if key == "" {
		return ""
	}

	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

func getEnvAsInteger(key string, defaultVal int) int {
	if key == "" {
		return 0
	}

	val, err := strconv.Atoi(os.Getenv(key))
	if val == 0 || err != nil {
		return defaultVal
	}

	return val
}
