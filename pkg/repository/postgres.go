package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	accountTable       = "account"
	interestsTable     = "interests"
	interestsUserTable = "interestsUser"
	likesTable         = "likes"
	premiumTable       = "premium"
	statusUserTable    = "statusUser"
	countryTable       = "country"
	cityTable          = "city"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// fmt.Println(cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	// db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
