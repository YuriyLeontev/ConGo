package repository

import (
	"congo"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type AccountsPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountsPostgres {
	return &AccountsPostgres{db: db}
}

func (r *AccountsPostgres) GetAll() ([]congo.Account, error) {
	var accounts []congo.Account

	query := fmt.Sprintf("SELECT * FROM %s", accountTable)
	err := r.db.Select(&accounts, query)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("GetAll", len(accounts))
	}

	return accounts, err
}

func (r *AccountsPostgres) Filter(filters []congo.Filter, limit int) ([]congo.Account, error) {
	var accounts []congo.Account
	var err error

	fmt.Println(limit)
	fmt.Println(filters)

	for _, filter := range filters {
		if strings.Compare(filter.Filter, "sex") == 0 {
			query := fmt.Sprintf("SELECT * FROM %s WHERE sex = '%s'", accountTable, filter.Parametr[0])
			err := r.db.Select(&accounts, query)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				if len(accounts) != 0 {
					fmt.Println("FilterSex", len(accounts))
				}
			}
		}
	}

	return accounts, err
}
