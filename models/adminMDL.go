package models

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/rafischer1/slr_capstone_go/db"
)

// Datum struct for the the psql table data
type Admin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAdmin varifies admin by matching password
func GetAdmin(password string) string {
	fmt.Println("In the get admin model", password)

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
		// gotta get all the fields!
		row.Scan(&admin.Username)
		entry = append(entry, admin)
		user = admin.Username
	}

	if err := row.Err(); err != nil {
		fmt.Println("err in admin model:", err)
		log.Fatal(err)
	}

	return user
}
