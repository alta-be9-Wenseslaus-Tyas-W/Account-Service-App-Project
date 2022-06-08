package main

import (
	"database/sql"
	"fmt"
	_config "project1/config"
	_controllTopUps "project1/controllers/top-up"
	_controllTransfers "project1/controllers/transfers"
	_controllUsers "project1/controllers/users"
	_entities "project1/entities"
)

var DBConn *sql.DB

func init() {
	DBConn = _config.ConnectionDB()
}

func main() {
	defer DBConn.Close()
	var ticket = false
	var id int
	for !ticket {
		var pilih int
		fmt.Println("Selamat Datang")
		fmt.Println("1. Log In")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Println("")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			fmt.Println("Masukkan Nomer Telpon")
			var telp string
			fmt.Scan(&telp)
			id = _controllUsers.GetIdUsersByTelp(DBConn, telp)
			if id == -1 {
				fmt.Println("Account tidak ditemukan, silahkan melakukan register terlebih dahulu")
			} else {
				ticket = true
			}
		case 2:
			newUser := _entities.Users{}
			fmt.Println("Masukkan Nama Lengkap")
			fmt.Scan(&newUser.NamaLengkap)
			fmt.Println("Masukkan Nick Name")
			fmt.Scan(&newUser.NickName)
			fmt.Println("Masukkan Nomer Telepon")
			fmt.Scan(&newUser.Telp)
			fmt.Println("Masukkan Password")
			var pass string
			fmt.Scan(&pass)
			_, err := _controllUsers.PostNewUser(DBConn, newUser)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Berhasil Register")
			}
			ticket = true
		}
	}
	for ticket {
		var pilih int
		fmt.Println("1. Read Account")
		fmt.Println("2. Update Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Top Up Saldo")
		fmt.Println("5. Transfer antar Account")
		fmt.Println("6. History Top Up")
		fmt.Println("7. History Transfer")
		fmt.Println("8. Lihat Account Orang Lain")
		fmt.Println("9. Exit")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
		case 2:
		case 3:
			fmt.Println("Apakah Anda Yakin Akan Menghapus Akun Anda?")
			fmt.Println("1. Ya. Saya akan menghapus akun milik saya")
			fmt.Println("2. Tidak. Syaa tidak yakin untuk menghapus akun milik saya")
			var pilih1 int
			fmt.Scan(&pilih1)
			if pilih1 == 1 {
				_controllUsers.DeleteUser(DBConn, id)
			}
		case 4:
			var nominal int
			fmt.Println("Masukkan Nominal Top Up:")
			fmt.Scan(&nominal)
			_, err := _controllTopUps.PostTopUp(DBConn, id, nominal)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Top Up Berhasil")
			}
		case 5:
			fmt.Print("Masukkan Nomor Penerima:")
			var telpPenerima string
			fmt.Scan(&telpPenerima)
			fmt.Print("\n")
			fmt.Print("Masukkan Nominal Transfer:")
			var nominal int
			fmt.Scan(&nominal)
			fmt.Print("\n")
			var idPenerima = _controllUsers.GetIdUsersByTelp(DBConn, telpPenerima)
			_, err := _controllTransfers.PostTransfer(DBConn, id, idPenerima, nominal)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Transfer Berhasil")
			}
		case 6:
		case 7:
		case 8:
		case 9:
			fmt.Println("Terimakasih Atas Kunjungannya")
			ticket = false
		}
	}
}
