package useraccounts

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_entities "project1/entities"
)

func GetAllUserAccounts(db *sql.DB) []_entities.UserAccount {
	results, err := db.Query("SELECT id_user_accounts, user_password FROM users_accounts")
	if err != nil {
		fmt.Println("error", err.Error())
	}

	var dataAll []_entities.UserAccount
	for results.Next() {
		var useraccounts _entities.UserAccount
		err := results.Scan(&useraccounts.IdUserAccounts, &useraccounts.Password)

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
	var enPass = GetMD5Hash(password)
	result, err := statement.Exec(&id, &enPass)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func GetUserPassword(db *sql.DB, id int) string {
	var query = "SELECT user_password FROM user_accounts WHERE id_user_accounts = ?"
	var pass string
	result := db.QueryRow(query, &id).Scan(&pass)
	if result != nil {
		if result == sql.ErrNoRows {
			return ""
		}
		return ""
	}
	return pass
}

func PutUserPassword(db *sql.DB, newPass string, id int) (int, error) {
	var query = "update user_accounts set user_password = ? where id_user_accounts = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	var pass = GetMD5Hash(newPass)
	result, err := statement.Exec(&pass, &id)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func DeleteUserPassword(db *sql.DB, id int) (int, error) {
	var query = "delete from user_accounts where id_user_accounts = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(&id)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
func GetMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}
