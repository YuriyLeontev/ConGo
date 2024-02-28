package repository

import (
	"congo"
	"fmt"

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
		fmt.Println("GetAll")
		fmt.Println(accounts[len(accounts)-1])
	}

	return accounts, err
}

func (r *AccountsPostgres) FilterSex(sex string) ([]congo.Account, error) {
	var accounts []congo.Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE sex = '%s'", accountTable, sex)
	err := r.db.Select(&accounts, query)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(accounts) != 0 {
			fmt.Println("FilterSex")
			fmt.Println(accounts[len(accounts)-1])
		} else {
			fmt.Println("FilterSex not found")
		}
	}

	return accounts, err
}
