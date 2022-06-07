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
	var pilih int
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("Masukkan Nomer Telpon")
		var telp string
		fmt.Scan(&telp)
		id := _controllUsers.GetIdUsersByTelp(DBConn, telp)
		fmt.Println(id)
	case 2:
		newUser := _entities.Users{}
		fmt.Println("Masukkan Nama")
		fmt.Scan(&newUser.NamaLengkap)
		fmt.Println("Masukkan Nick Name")
		fmt.Scan(&newUser.NickName)
		fmt.Println("Masukkan Nomer Telepon")
		fmt.Scan(&newUser.Telp)
		fmt.Println("Masukkan Password")
		var pass string
		fmt.Scan(&pass)

	}

}
