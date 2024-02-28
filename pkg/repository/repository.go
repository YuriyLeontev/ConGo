package repository

import (
	"congo"

	"github.com/jmoiron/sqlx"
)

type AccountsList interface {
	GetAll() ([]congo.Account, error)
	FilterSex(string) ([]congo.Account, error)
}

type Repository struct {
	AccountsList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountsList: NewAccountPostgres(db),
	}
}
