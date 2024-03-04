package repository

import (
	"congo"

	"github.com/jmoiron/sqlx"
)

type AccountsList interface {
	GetAll() ([]congo.Account, error)
	Filter([]congo.Filter, int) ([]congo.Account, error)
}

type Repository struct {
	AccountsList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountsList: NewAccountPostgres(db),
	}
}
