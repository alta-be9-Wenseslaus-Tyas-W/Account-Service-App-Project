package transfers

import (
	"database/sql"
	"fmt"
	_controllUsers "project1/controllers/users"
	_entities "project1/entities"
)

func GetAllTransfers(db *sql.DB) ([]_entities.Transfers, error) {
	var query = "select id_transfers, id_user_pemberi, id_user_penerima, nominal , sisa_saldo from transfers"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return []_entities.Transfers{}, errPrepare
	}
	result, err := statement.Query()
	if err != nil {
		return []_entities.Transfers{}, err
	}
	var historyTransfers = []_entities.Transfers{}
	for result.Next() {
		var transfer = _entities.Transfers{}
		err := result.Scan(&transfer.IdTransfers, &transfer.IdUserPenerima, &transfer.IdUserPengirim, &transfer.Nominal, &transfer.SisaSaldo)
		if err != nil {
			return []_entities.Transfers{}, err
		}
		historyTransfers = append(historyTransfers, transfer)
	}
	return historyTransfers, nil
}

func PostTransfer(db *sql.DB, idPemberi int, idPenerima int, nominal int) (int, error) {
	var query = "insert into transfers(id_user_pengirim, id_user_penerima, nominal , sisa_saldo) values (?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldoPemberi, errSaldo := _controllUsers.GetUserSaldo(db, idPemberi)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var sisaSaldo int
	if saldoPemberi > nominal && saldoPemberi > 10000 {
		sisaSaldo = saldoPemberi - nominal
		_, errPostKurang := _controllUsers.PostKurangSaldo(db, idPemberi, nominal)
		if errPostKurang != nil {
			return 0, errPostKurang
		}
	} else {
		fmt.Println("Saldo tidak mencukupi")
		sisaSaldo = saldoPemberi
		nominal = 0
	}
	result, err := statement.Exec(&idPemberi, &idPenerima, &nominal, &sisaSaldo)
	_, errPostTambah := _controllUsers.PostTambahSaldo(db, idPenerima, nominal)
	if errPostTambah != nil {
		return 0, errPostTambah
	}
	if err != nil {
		return 0, err
	} else {
		rowTopUp, _ := result.RowsAffected()
		return int(rowTopUp), nil
	}
}

func GetHistoryTransferById(db *sql.DB, idUser int) ([]_entities.HistoryTransfer, error) {
	var query = "select us.nama_lengkap, ur.nama_lengkap, tf.nominal , tf.sisa_saldo from transfers tf inner join users us on tf.id_user_pengirim = us.id_user inner join users ur on tf.id_user_penerima = ur.id_user where us.id_user = ? order by tf.id_transfers desc"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return []_entities.HistoryTransfer{}, errPrepare
	}
	result, err := statement.Query(&idUser)
	if err != nil {
		return []_entities.HistoryTransfer{}, err
	}
	var historyTransfers = []_entities.HistoryTransfer{}
	for result.Next() {
		var transfer = _entities.HistoryTransfer{}
		err := result.Scan(&transfer.NamaPengirim, &transfer.NamaPenerima, &transfer.Nominal, &transfer.SisaSaldo)
		if err != nil {
			return []_entities.HistoryTransfer{}, err
		}
		historyTransfers = append(historyTransfers, transfer)
	}
	return historyTransfers, nil
}
