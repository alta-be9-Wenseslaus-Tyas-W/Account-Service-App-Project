package topup

import (
	"database/sql"
	"fmt"
	_controllUsers "project1/controllers/users"
	_entities "project1/entities"
)

func GetAllTopUp(db *sql.DB) ([]_entities.TopUp, error) {
	var query = "select id_top_ups , id_user , nominal from top_ups"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return []_entities.TopUp{}, errPrepare
	}
	result, err := statement.Query()
	defer db.Close()
	if err != nil {
		return []_entities.TopUp{}, err
	}
	var historyTopUp = []_entities.TopUp{}
	for result.Next() {
		var topup = _entities.TopUp{}
		err := result.Scan(&topup.IdTopUps, &topup.IdUser, &topup.Nominal)
		if err != nil {
			return []_entities.TopUp{}, err
		}
		historyTopUp = append(historyTopUp, topup)
	}
	return historyTopUp, nil
}

func PostTopUp(db *sql.DB, idUser int, nominal int) (int, error) {
	var query = "insert into top_ups (id_user,nominal) values (?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	_controllUsers.PostTambahSaldo(db, idUser, nominal)
	result, err := statement.Exec(idUser, nominal)
	defer db.Close()
	if err != nil {
		return 0, err
	} else {
		rowTopUp, _ := result.RowsAffected()
		fmt.Println("Saldo anda bertambah")
		return int(rowTopUp), nil
	}
}
