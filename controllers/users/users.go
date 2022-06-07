package users

import (
	"database/sql"
	"fmt"
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
