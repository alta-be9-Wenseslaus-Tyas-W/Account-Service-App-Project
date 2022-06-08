package main

import (
	"database/sql"
	"fmt"
	_config "project1/config"
	_controllUsers "project1/controllers/users"
	_entities "project1/entities"
)

var DBConn *sql.DB

func init() {
	DBConn = _config.ConnectionDB()
}

func main() {
	defer DBConn.Close()
	fmt.Println("Selamat Datang")
	fmt.Println("1. Log In")
	fmt.Println("2. Register")
	// tester
	fmt.Println("3. Read Account")
	fmt.Println("5. Delete Account")
	var pilih int
	var id int
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Masukkan Nomer Telpon")
		var telp string
		fmt.Scan(&telp)
		id = _controllUsers.GetIdUsersByTelp(DBConn, telp)
		fmt.Println(id)

		fmt.Println("Apakah Anda Yakin Akan Menghapus Akun Anda?")
		fmt.Println("1. Ya. Saya akan menghapus akun milik saya")
		fmt.Println("2. Tidak. Syaa tidak yakin untuk menghapus akun milik saya")
		var pilih1 int
		fmt.Scan(&pilih1)
		if pilih1 == 1 {
			_controllUsers.DeleteUser(DBConn, id)
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
		row, err := _controllUsers.PostNewUser(DBConn, newUser)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Berhasil Register")
			fmt.Println("Baris Bertambah", row)
		}

		// tester
		// case 3:
		// 	results := _controllUsers.PostNewUser(DBConn)
		// 	for _, v := range results {
		// 		fmt.Println("ID:", v.id_user, "Nama Lengkap:", v.nama_lengkap, "Nick Name:", v.nick_name, "Nomer Telepon:", v.nomer_telp, "Saldo:", v.saldo)
		// 	}

	case 5:
		fmt.Println("Apakah Anda Yakin Akan Menghapus Akun Anda?")
		fmt.Println("1. Ya. Saya akan menghapus akun milik saya")
		fmt.Println("2. Tidak. Syaa tidak yakin untuk menghapus akun milik saya")
		var pilih1 int
		fmt.Scan(&pilih1)
		if pilih1 == 1 {
			_controllUsers.DeleteUser(DBConn, id)
		}
	}
}
