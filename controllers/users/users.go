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

func GetUserSaldo(db *sql.DB, idUser int) int {
	var query = "SELECT saldo FROM users WHERE id_user = ?"
	var saldo int
	result := db.QueryRow(query, idUser).Scan(&saldo)
	defer db.Close()
	if result != nil {
		if result == sql.ErrNoRows {
			fmt.Println(result.Error())
			return -1
		}
		fmt.Println(result.Error())
		return -1
	}
	return saldo
}

func PostTambahSaldo(db *sql.DB, idUser int, nominal int) {
	var query = "update users set saldo = (?) where id_user = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		fmt.Println("error", errPrepare.Error())
	}
	var newSaldo = nominal + GetUserSaldo(db, idUser)
	result, err := statement.Exec(newSaldo, idUser)
	defer db.Close()
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		row, _ := result.RowsAffected()
		fmt.Println(row)
	}
}
