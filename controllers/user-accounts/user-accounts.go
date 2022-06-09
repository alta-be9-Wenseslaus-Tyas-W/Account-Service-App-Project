package useraccounts

import (
	"crypto/aes"
	"database/sql"
	"encoding/hex"
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

func GetUserPassword(db *sql.DB, id int) string {
	var query = "SELECT user_password FROM user_accounts WHERE id_user_accounts = ?"
	var pass string
	result := db.QueryRow(query, &id).Scan(&pass)
	if result != nil {
		if result == sql.ErrNoRows {
			fmt.Println(result.Error())
			return ""
		}
		fmt.Println(result.Error())
		return ""
	}
	return pass
}

func CheckPassword(inputPass string, dataPass string) bool {
	return inputPass == dataPass
}

func EncryptPassword(key string, message string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(message))
	c.Encrypt(msgByte, []byte(message))
	return hex.EncodeToString(msgByte)
}

func DecryptPassword(key string, message string) string {
	txt, _ := hex.DecodeString(message)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(txt))
	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}
