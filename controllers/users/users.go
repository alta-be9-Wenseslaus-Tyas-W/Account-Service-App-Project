package users

import (
	"database/sql"
	"fmt"
	_entities "project1/entities"
)

func GetAllUsers() {

}

func GetIdUsersByTelp(db *sql.DB, telp string) int {
	var query = "SELECT id_user FROM users WHERE telp = ?"
	var id int
	result := db.QueryRow(query, telp).Scan(&id)
	if result != nil {
		if result == sql.ErrNoRows {
			fmt.Println(result.Error())
			return -1
		}
		fmt.Println(result.Error())
		return -1
	}
	return id
}

func PostNewUser(db *sql.DB, newUser _entities.Users) (int, error) {
	var queryuser = (`INSERT INTO users (nama_lengkap, nick_name, telp, saldo) VALUES (?, ?, ?, ?)`)
	statement, errPrepare := db.Prepare(queryuser)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(newUser.NamaLengkap, newUser.NickName, newUser.Telp, 0)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
