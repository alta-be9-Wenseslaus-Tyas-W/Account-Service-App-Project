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
	var idAccount int
	for !ticket {
		var pilih int
		fmt.Println("Selamat Datang")
		fmt.Println("1. Log In")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		fmt.Print("\n")
		if pilih == 1 {
			fmt.Print("Masukkan Nomer Telpon: ")
			var telp string
			fmt.Scan(&telp)
			fmt.Print("\n")
			idAccount = _controllUsers.GetIdUsersByTelp(DBConn, telp)
			if idAccount == -1 {
				fmt.Println("Account tidak ditemukan, silahkan melakukan register terlebih dahulu")
			} else {
				fmt.Println("Account terdaftar")
				ticket = true
			}
			fmt.Print("\n")
		} else if pilih == 2 {
			fmt.Println("Silahkan isi data yang tersedia")
			newUser := _entities.Users{}
			fmt.Print("Masukkan Nama Lengkap: ")
			fmt.Scan(&newUser.NamaLengkap)
			fmt.Print("\n")
			fmt.Print("Masukkan Nick Name: ")
			fmt.Scan(&newUser.NickName)
			fmt.Print("\n")
			fmt.Println("Masukkan Nomer Telepon: ")
			fmt.Scan(&newUser.Telp)
			fmt.Print("\n")
			fmt.Print("Masukkan Password: ")
			var pass string
			fmt.Scan(&pass)
			fmt.Print("\n")
			_, err := _controllUsers.PostNewUser(DBConn, newUser)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Berhasil Register")
			}
		} else {
			fmt.Println("Terimakasih Atas Kunjungannya")
			break
		}
	}
	for ticket {
		var pilih int
		defer DBConn.Begin()
		fmt.Println("SELAMAT DATANG")
		fmt.Println("Silahkan Pilih Fitur yang Tersedia")
		fmt.Println("1. Read Account")
		fmt.Println("2. Update Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Top Up Saldo")
		fmt.Println("5. Transfer antar Account")
		fmt.Println("6. History Top Up")
		fmt.Println("7. History Transfer")
		fmt.Println("8. Lihat Account Orang Lain")
		fmt.Println("9. Exit")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilih)
		fmt.Print("\n")
		switch pilih {
		case 1:
			result := _controllUsers.ReadUserInfo(DBConn, idAccount)
			fmt.Println("Nama: ", result.NamaLengkap)
			fmt.Println("Nick Name: ", result.NickName)
			fmt.Println("Nomer Telpon: ", result.Telp)
			fmt.Println("Saldo: ", result.Saldo)
			fmt.Println()
		case 2:
		case 3:
			fmt.Println("Delete Account Anda")
			fmt.Println("Apakah Anda Yakin Akan Menghapus Account Anda?")
			fmt.Println("1. Ya. Saya akan menghapus account milik saya")
			fmt.Println("2. Tidak. Saya tidak yakin untuk menghapus account milik saya")
			var pilih1 int
			fmt.Print("Pilihan saya:")
			fmt.Scan(&pilih1)
			fmt.Print("\n")
			if pilih1 == 1 {
				_controllUsers.DeleteUser(DBConn, idAccount)
				ticket = false
			} else {
				fmt.Println("Account Tidak Dihapus oleh User")
			}
			fmt.Print("\n")
		case 4:
			var nominal int
			fmt.Print("Silahkan Masukkan Nominal Top Up: ")
			fmt.Scan(&nominal)
			fmt.Print("\n")
			_, err := _controllTopUps.PostTopUp(DBConn, idAccount, nominal)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Top Up Berhasil")
			}
			fmt.Print("\n")
		case 5:
			fmt.Print("Masukkan Nomor Penerima: ")
			var telpPenerima string
			fmt.Scan(&telpPenerima)
			fmt.Print("\n")
			fmt.Print("Masukkan Nominal Transfer: ")
			var nominal int
			fmt.Scan(&nominal)
			fmt.Print("\n")
			var idPenerima = _controllUsers.GetIdUsersByTelp(DBConn, telpPenerima)
			_, err := _controllTransfers.PostTransfer(DBConn, idAccount, idPenerima, nominal)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Transfer Berhasil")
			}
			fmt.Print("\n")
		case 6:
			fmt.Println("History Top Up Account Anda:")
			result, err := _controllTopUps.GetHistoryTopUpById(DBConn, idAccount)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				for _, v := range result {
					fmt.Print("Nominal Top up: ", v.Nominal, "\n")
				}
			}
			fmt.Print("\n")
		case 7:
			fmt.Println("History Transfer Account Anda:")
			result, err := _controllTransfers.GetHistoryTransferById(DBConn, idAccount)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				for _, v := range result {
					fmt.Printf("Nama Pengirim: %s \t Nama Penerima: %s \n", v.NamaPengirim, v.NamaPenerima)
					fmt.Printf("Nominal: %d \t Sisa Saldo: %d \n", v.Nominal, v.SisaSaldo)
				}
			}
			fmt.Print("\n")
		case 8:
			fmt.Print("Masukkan Nomer Telpon Account lain: ")
			var telp string
			fmt.Scan(&telp)
			fmt.Print("\n")
			idLain := _controllUsers.GetIdUsersByTelp(DBConn, telp)
			result := _controllUsers.ReadUserInfo(DBConn, idLain)
			fmt.Println("Nama: ", result.NamaLengkap)
			fmt.Println("Nick Name: ", result.NickName)
			fmt.Println("Nomer Telpon: ", result.Telp)
			fmt.Println()
		case 9:
			fmt.Println("Terimakasih Atas Kunjungannya")
			ticket = false
		}
	}
}
