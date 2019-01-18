package models

import (
	"database/sql"
	"log"

	d "github.com/rafischer1/slr_capstone_go/db"
)

// Datum struct for the the psql table data
type Datum struct {
	ID         int     `json:"id"`
	Msg        string  `json:"msg"`
	WindMPH    float64 `json:"windmph"`
	WindDir    string  `json:"winddir"`
	SeaLevelFt float64 `json:"sealevelft"`
	CreatedAt  string  `json:"createdat"`
}

// GetAllData model function
func GetAllData() []Datum {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM data")

	defer rows.Close()

	var events []Datum

	for rows.Next() {
		event := Datum{}

		// gotta get all the fields!
		rows.Scan(&event.ID, &event.Msg, &event.WindMPH, &event.WindDir, &event.SeaLevelFt, &event.CreatedAt)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return events
}

// PostData will hopefully be the big payoff for trying out the SMS service
func PostData(Msg string, WindMPH float64, WindDir string, SeaLevelFt float64) error {
	db, err := sql.Open("postgres", d.ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	event := Datum{}
	var entry []Datum

	//Create
	errTwo := db.QueryRow(`INSERT INTO data(msg, windmph, winddir, sealevelft) VALUES($1, $2, $3, $4) RETURNING *`, Msg, WindMPH, WindDir, SeaLevelFt).Scan(&event.ID, &event.Msg, &event.WindMPH, &event.WindDir, &event.SeaLevelFt, &event.CreatedAt)
	entry = append(entry, event)
	if errTwo != nil {
		return errTwo
	}

	return errTwo
}
