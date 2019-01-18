package models

import (
	"database/sql"
	"fmt"
	"log"

	d "github.com/rafischer1/slr_capstone_go/db"
)

// Subscriber struct for the the psql table subscribers
type Subscriber struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

// GetAllSubs model function
func GetAllSubs() []Subscriber {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("in the get all sub db:", db)

	defer db.Close()
	rows, err := db.Query("SELECT * FROM subscribers")

	defer rows.Close()

	var subscriptions []Subscriber

	for rows.Next() {
		subscription := Subscriber{}

		// gotta get all the fields!
		rows.Scan(&subscription.ID, &subscription.Phone, &subscription.Location)
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return subscriptions
}

// PostSub function
func PostSub(Phone string, Location string) error {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	subscription := Subscriber{}
	var entry []Subscriber
	//Create
	errTwo := db.QueryRow(`INSERT INTO subscribers(phone, location) VALUES($1, $2) RETURNING *`, Phone, Location).Scan(&subscription.ID, &subscription.Phone, &subscription.Location)
	entry = append(entry, subscription)
	if errTwo != nil {
		return errTwo
	}

	return errTwo
}

// DeleteSub Model function
func DeleteSub(phone string) (string, error) {
	fmt.Println("In the delete model", phone)

	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row, err := db.Query(`Delete FROM subscribers where phone = $1`, phone)
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return phone, err
}
