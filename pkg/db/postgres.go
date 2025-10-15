package db

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBname   string
	Port     string
}

func GetConfig(log *logrus.Logger) (*Config, error) {
	log.Println("Getting config from env")

	cfg := Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		DBname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	return &cfg, nil
}

func DBInit(log *logrus.Logger, cfg *Config) (db *sql.DB, err error) {
	connStr := GetConnStr(log, cfg)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to init db, error: %v", err)
		return nil, err
	}
	return db, nil
}

func GetConnStr(log *logrus.Logger, cfg *Config) string {
	log.Info("Getting connStr")

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
