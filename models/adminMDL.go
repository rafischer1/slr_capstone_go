package models

import (
	"database/sql"
	"log"

	d "github.com/rafischer1/slr_capstone_go/db"
)

// Admin struct for the the psql table `admin``
type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAdmin varifies admin by matching password and returning a username or an error
func GetAdmin(password string) (string, error) {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row, err := db.Query(`SELECT username FROM admin WHERE password=$1`, password)
	defer row.Close()

	var entry []Admin
	var user string

	for row.Next() {
		admin := Admin{}
		row.Scan(&admin.Username)
		entry = append(entry, admin)
		user = admin.Username
	}

	if err := row.Err(); err != nil {
		log.Fatal(err)
	}

	return user, err
}
