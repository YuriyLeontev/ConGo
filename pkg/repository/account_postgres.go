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
	var request_join string = ""

	for _, filter := range filters {
		if strings.Compare(filter.Filter, "sex") == 0 {
			request_filter = append(request_filter, fmt.Sprintf("sex = '%s'", filter.Parametr[0]))
		} else if strings.Compare(filter.Filter, "email") == 0 {
			if filter.Method == "domain" {
				request_filter = append(request_filter, " email ILIKE  '%"+fmt.Sprintf("%s' ", filter.Parametr[0]))
			} else if filter.Method == "lt" {
				request_filter = append(request_filter, fmt.Sprintf(" email < '%s' ", filter.Parametr[0]))
			} else if filter.Method == "gt" {
				request_filter = append(request_filter, fmt.Sprintf(" email > '%s' ", filter.Parametr[0]))
			}
		} else if strings.Compare(filter.Filter, "status") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf(" status_user = '%s' ", filter.Parametr[0]))
			} else if filter.Method == "neq" {
				request_filter = append(request_filter, fmt.Sprintf(" status_user <> '%s' ", filter.Parametr[0]))
			}
		} else if strings.Compare(filter.Filter, "fname") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf(" fname = '%s' ", filter.Parametr[0]))
			} else if filter.Method == "any" {
				names := "('" + strings.Join(filter.Parametr, "', '") + "')"
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
		} else if strings.Compare(filter.Filter, "sname") == 0 {
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
		} else if strings.Compare(filter.Filter, "phone") == 0 {
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
		} else if strings.Compare(filter.Filter, "country") == 0 {
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
		} else if strings.Compare(filter.Filter, "city") == 0 {
			if filter.Method == "eq" {
				request_filter = append(request_filter, fmt.Sprintf("city = '%s'", filter.Parametr[0]))
			} else if filter.Method == "any" {
				cities := "('" + strings.Join(filter.Parametr, "', '") + "')"
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
		} else if strings.Compare(filter.Filter, "birth") == 0 {
			if filter.Method == "lt" {
				request_filter = append(request_filter, fmt.Sprintf(" birth < '%s'", filter.Parametr[0]))
			} else if filter.Method == "gt" {
				request_filter = append(request_filter, fmt.Sprintf(" birth > '%s' ", filter.Parametr[0]))
			} else if filter.Method == "year" {
				request_filter = append(request_filter, fmt.Sprintf(" birth BETWEEN '01.01.%s' AND '12.31.%s' ", filter.Parametr[0], filter.Parametr[0]))
			}
		} else if strings.Compare(filter.Filter, "interests") == 0 {
			interest := "('" + strings.Join(filter.Parametr, "', '") + "')"
			if filter.Method == "contains" {
				//SELECT acc.* FROM accounts acc INNER JOIN (SELECT account_id FROM accounts_interest WHERE interest IN ('Любовь', 'Солнце')
				//          GROUP BY account_id HAVING count(account_id)>1) i on i.account_id = acc.id

				count_param := len(filter.Parametr) - 1
				req := fmt.Sprintf("(SELECT account_id FROM accounts_interest WHERE interest IN %s GROUP BY account_id HAVING count(account_id)>%d)", interest, count_param)

				request_join = fmt.Sprintf("INNER JOIN %s i on i.account_id = acc.id", req)

			} else if filter.Method == "any" {
				request_join = fmt.Sprintf("INNER JOIN %s i on i.account_id = acc.id", interestsTable)
				request_filter = append(request_filter, fmt.Sprintf(" i.interest IN %s", interest))
			}
		} else if strings.Compare(filter.Filter, "interests") == 0 {
			if filter.Method == "contains" {
				// SELECT acc.* FROM accounts acc INNER JOIN
				// (select user_id from accounts_like where account_id in('5208','6478') group by user_id having count(*) = 2) i on i.user_id = acc.id

				// accounts := "('" + strings.Join(filter.Parametr, "', '") + "')"
				// count_param := len(filter.Parametr)
			}
		}
		// TO-DO premium (now, null)
	}

	whereQuery := request_join
	if len(request_filter) != 0 {
		whereQuery += " WHERE "
	}

	whereQuery += strings.Join(request_filter, " AND ")

	query := fmt.Sprintf("SELECT acc.* FROM %s acc %s LIMIT %d", accountTable, whereQuery, limit)
	fmt.Println(limit)
	fmt.Println(query)

	err := r.db.Select(&accounts, query)

	if err != nil {
		return nil, err
	}

	if len(accounts) != 0 {
		fmt.Println("Accounts", len(accounts))
	}

	return accounts, err
}
