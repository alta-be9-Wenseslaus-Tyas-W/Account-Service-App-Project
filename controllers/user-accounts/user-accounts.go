package useraccounts

import (
	"database/sql"
	"fmt"
	_entities "project1/entities"
)

func GetAllUserAccounts(db *sql.DB) []_entities.UserAccount {
	results, err := db.Query("SELECT password FROM Users-Accounts")
	if err != nil {
		fmt.Println("error", err.Error())
	}

	var dataAll []_entities.UserAccount
	for results.Next() {
		var useraccounts _entities.UserAccount
		err := results.Scan(&useraccounts.Password)

		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		dataAll = append(dataAll, useraccounts)
	}

	return dataAll
}

func PostNewPassword(db *sql.DB, id int, password string) (int, error) {
	var pass = (`INSERT INTO user_accounts (id_user_accounts, user_password) VALUES (?, ?)`)
	statement, errPrepare := db.Prepare(pass)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(&id, &password)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}

}
