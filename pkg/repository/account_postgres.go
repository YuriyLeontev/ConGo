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

var operations = map[string]string{
	"eq":     "=",
	"neq":    "<>",
	"lt":     "<",
	"gt":     ">",
	"domain": "ILIKE",
	"starts": "ILIKE",
	"code":   "ILIKE",
	"any":    "IN",
	"null":   "IS",
	"year":   "BETWEEN",
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

func (r *AccountsPostgres) addFilterRequest(filters []congo.Filter) []string {
	var request_filter []string
	for _, filter := range filters {
		request_filter = append(request_filter, fmt.Sprintf("%s %s %s", filter.Filter, operations[filter.Method], filter.Parametr))
	}
	return request_filter
}

func (r *AccountsPostgres) Filter(filters []congo.Filter, limit int) ([]congo.Account, error) {
	var accounts []congo.Account
	var request_filter []string
	var request_join string = ""

	for i, filter := range filters {
		switch filter.Filter {
		case "email":
			switch filter.Method {
			case "domain":
				filters[i].Parametr = "'%" + filter.Parametr + "'"
			}
		case "status":
			filters[i].Filter = "status_user"
		}

		switch filter.Method {
		case "null":
			switch filter.Parametr {
			case "0":
				filters[i].Parametr = "NULL"
			case "1":
				filters[i].Parametr = "NOT NULL"
			}
		case "any":
			params := strings.Join(strings.Split(filter.Parametr, ","), "', '")
			filters[i].Parametr = fmt.Sprintf("('%s')", params)
		case "code":
			filters[i].Parametr = "%(" + filter.Parametr + ")%"
		case "year":
			filters[i].Parametr = fmt.Sprintf("01.01.%s' AND '12.31.%s", filter.Parametr, filter.Parametr)

		case "eq", "neq", "lt", "gt":
			filters[i].Parametr = fmt.Sprintf("'%s'", filters[i].Parametr)
		}
	}

	request_filter = r.addFilterRequest(filters)

	for _, filter := range filters {
		switch filter.Filter {
		case "interests":
			interest := "('" + filter.Parametr + "')"
			switch filter.Method {
			case "contains":
				//SELECT acc.* FROM accounts acc INNER JOIN (SELECT account_id FROM accounts_interest WHERE interest IN ('Любовь', 'Солнце')
				//          GROUP BY account_id HAVING count(account_id)>1) i on i.account_id = acc.id

				count_param := len(filter.Parametr) - 1
				req := fmt.Sprintf("(SELECT account_id FROM accounts_interest WHERE interest IN %s GROUP BY account_id HAVING count(account_id)>%d)", interest, count_param)

				request_join = fmt.Sprintf("INNER JOIN %s i on i.account_id = acc.id", req)
			case "any":
				request_join = fmt.Sprintf("INNER JOIN %s i on i.account_id = acc.id", interestsTable)
				request_filter = append(request_filter, fmt.Sprintf(" i.interest IN %s", interest))
			}

		case "likes":
			switch filter.Method {
			case "contains":
				// SELECT acc.* FROM accounts acc INNER JOIN
				// (select user_id from accounts_like where account_id in('5208','6478') group by user_id having count(*) = 2) i on i.user_id = acc.id

				// accounts := "('" + strings.Join(filter.Parametr, "', '") + "')"
				// count_param := len(filter.Parametr)
			}
		case "premium":
			// TO-DO premium (now, null)
		}
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
