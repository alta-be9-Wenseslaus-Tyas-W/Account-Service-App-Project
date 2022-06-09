package users

import (
	"database/sql"
	_entities "project1/entities"
)

func GetAllUsers() {

}

func GetIdUsersByTelp(db *sql.DB, telp string) int {
	var query = "SELECT id_user FROM users WHERE telp = ?"
	var id int
	result := db.QueryRow(query, &telp).Scan(&id)
	if result != nil {
		if result == sql.ErrNoRows {
			return -1
		}
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
	result, err := statement.Exec(&newUser.NamaLengkap, &newUser.NickName, &newUser.Telp, 0)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func GetUserSaldo(db *sql.DB, idUser int) (int, error) {
	var query = "SELECT saldo FROM users WHERE id_user = ?"
	var saldo int
	result := db.QueryRow(query, &idUser).Scan(&saldo)
	if result != nil {
		if result == sql.ErrNoRows {
			return -1, result
		}
		return -1, result
	}
	return saldo, nil
}

func PostTambahSaldo(db *sql.DB, idUser int, nominal int) (int, error) {
	var query = "update users set saldo = (?) where id_user = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldo, errSaldo := GetUserSaldo(db, idUser)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var newSaldo = nominal + saldo
	result, err := statement.Exec(&newSaldo, &idUser)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func PostKurangSaldo(db *sql.DB, idUser int, nominal int) (int, error) {
	var query = "update users set saldo = (?) where id_user = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldo, errSaldo := GetUserSaldo(db, idUser)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var newSaldo = saldo - nominal
	result, err := statement.Exec(&newSaldo, &idUser)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func DeleteUser(db *sql.DB, id int) (int, error) {
	var delete = "DELETE from users WHERE id_user = ? "
	statment, errPrepare := db.Prepare((delete))
	if errPrepare != nil {
		return 0, errPrepare
	}
	var result, err = statment.Exec(&id)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func ReadUserInfo(db *sql.DB, id int) _entities.Users {
	results := db.QueryRow("SELECT id_user, nama_lengkap, nick_name, telp, saldo from users where id_user = ?", &id)

	var dataUser _entities.Users
	err := results.Scan(&dataUser.IdUser, &dataUser.NamaLengkap, &dataUser.NickName, &dataUser.Telp, &dataUser.Saldo)

	if err != nil {
		return _entities.Users{}
	}

	return dataUser
}

func PutDataUser(db *sql.DB, query string, id int) (int, error) {
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
