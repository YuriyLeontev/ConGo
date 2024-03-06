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
				request_filter = append(request_filter, fmt.Sprintf(" email < '%s' ", filter.Parametr[0]))
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
					// Указана фамилия
					request_filter = append(request_filter, "sname IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указана фамилия
					request_filter = append(request_filter, "sname IS NOT NULL")
				}
			}
		}
		if strings.Compare(filter.Filter, "phone") == 0 {
			if filter.Method == "code" {
				request_filter = append(request_filter, " phone ILIKE '%("+filter.Parametr[0]+")%'")
			} else if filter.Method == "null" {
				if filter.Parametr[0] == "0" {
					// Указан номер
					request_filter = append(request_filter, "phone IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указан номер
					request_filter = append(request_filter, "phone IS NOT NULL")
				}
			}
		}
		if strings.Compare(filter.Filter, "country") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf("country = '%s'", filter.Parametr[0]))
			} else if filter.Method == "null" {
				if filter.Parametr[0] == "0" {
					// Указана страна
					request_filter = append(request_filter, "country IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указана страна
					request_filter = append(request_filter, "country IS NOT NULL")
				}
			}
		}
		if strings.Compare(filter.Filter, "city") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf("city = '%s'", filter.Parametr[0]))
			} else if filter.Method == "any" {
				cities := "("
				for i, city := range filter.Parametr {
					if i != 0 {
						cities += ", "
					}
					cities += fmt.Sprintf("'%s'", city)
				}
				cities += ")"
				request_filter = append(request_filter, fmt.Sprintf(" city IN %s", cities))
			} else if filter.Method == "null" {
				if filter.Parametr[0] == "0" {
					// Указан город
					request_filter = append(request_filter, "city IS NULL")
				} else if filter.Parametr[0] == "1" {
					// Не указан город
					request_filter = append(request_filter, "city IS NOT NULL")
				}
			}
		}
		if strings.Compare(filter.Filter, "birth") == 0 {
			if filter.Method == "lt" {
				request_filter = append(request_filter, fmt.Sprintf(" birth < '%s'", filter.Parametr[0]))
			} else if filter.Method == "gt" {
				request_filter = append(request_filter, fmt.Sprintf(" birth > '%s' ", filter.Parametr[0]))
			} else if filter.Method == "year" {
				request_filter = append(request_filter, fmt.Sprintf(" birth BETWEEN '01.01.%s' AND '12.31.%s' ", filter.Parametr[0], filter.Parametr[0]))
			}
		}

		// TO-DO interests (contains, any)
		// TO-DO likes (contains)
		// TO-DO premium (now, null)
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
