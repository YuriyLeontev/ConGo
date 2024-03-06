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
	var request_filter []string

	for _, filter := range filters {
		if strings.Compare(filter.Filter, "sex") == 0 {
			request_filter = append(request_filter, fmt.Sprintf("sex = '%s'", filter.Parametr[0]))
		}
		if strings.Compare(filter.Filter, "email") == 0 {
			if filter.Method == "domain" {
				request_filter = append(request_filter, " email ILIKE  '%"+fmt.Sprintf("%s' ", filter.Parametr[0]))
			} else if filter.Method == "lt" {
				request_filter = append(request_filter, fmt.Sprint(" email < '%s' ", filter.Parametr[0]))
			} else if filter.Method == "gt" {
				request_filter = append(request_filter, fmt.Sprintf(" email > '%s' ", filter.Parametr[0]))
			}
		}
		if strings.Compare(filter.Filter, "status") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf(" status_user = '%s' ", filter.Parametr[0]))
			} else if filter.Method == "neq" {
				request_filter = append(request_filter, fmt.Sprintf(" status_user <> '%s' ", filter.Parametr[0]))
			}
		}
		if strings.Compare(filter.Filter, "fname") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf(" fname = '%s' ", filter.Parametr[0]))
			} else if filter.Method == "any" {
				names := "("
				for i, name := range filter.Parametr {
					if i != 0 {
						names += ", "
					}
					names += fmt.Sprintf("'%s'", name)
				}
				names += ")"
				request_filter = append(request_filter, fmt.Sprintf(" fname IN %s", names))
			} else if filter.Method == "null" {
				if filter.Parametr[0] == "0" {
					// Указано имя
					request_filter = append(request_filter, "fname IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указано имя
					request_filter = append(request_filter, "fname IS NOT NULL")
				}
			}
		}
		if strings.Compare(filter.Filter, "sname") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf(" sname = '%s' ", filter.Parametr[0]))
			} else if filter.Method == "starts" {
				request_filter = append(request_filter, fmt.Sprintf("sname ILIKE '%s", filter.Parametr[0])+"%'")
			} else if filter.Method == "null" {
				if filter.Parametr[0] == "0" {
					// Указано имя
					request_filter = append(request_filter, "sname IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указано имя
					request_filter = append(request_filter, "sname IS NOT NULL")
				}
			}
		}
	}

	where := ""
	if len(request_filter) != 0 {
		where = "WHERE "
	}

	for i, filter := range request_filter {
		if i != 0 {
			where += " AND "
		}
		where += filter
	}

	fmt.Println(limit)
	fmt.Println(filters)
	fmt.Println(where)

	query := fmt.Sprintf("SELECT * FROM %s %s LIMIT %d", accountTable, where, limit)
	err := r.db.Select(&accounts, query)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		if len(accounts) != 0 {
			fmt.Println("Accounts", len(accounts))
		}
	}

	return accounts, err
}
