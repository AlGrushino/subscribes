package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
}

func GetConfig() (*Config, error) {
	cfg := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		DBname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	return &cfg, nil
}

func DBInit(cfg *Config) (db *sql.DB, err error) {
	// connStr := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s",
	// 	cfg.User,
	// 	cfg.Password,
	// 	cfg.Host,
	// 	cfg.Port,
	// 	cfg.DBname,
	// )

	connStr := GetConnStr(cfg)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetConnStr(cfg *Config) string {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBname,
	)
	return connStr
}
