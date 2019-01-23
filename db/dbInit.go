package db

import (
	"os"

	"github.com/subosito/gotenv"
)

// ConnStr package initializes the db connection (check production vs dev build)
var ConnStr = Init()

//FOR DEVELEOPMENT ONLY !!!!

// Init initializes the database using .env vars
// func Init() string {
// 	gotenv.Load()
// 	dbname := os.Getenv("DBNAME")
// 	dbuser := os.Getenv("DBUSER")
// 	ConnStr := fmt.Sprintf("user=%[1]v "+
// 		"dbname=%[2]v sslmode=disable", dbuser, dbname)
// 	return ConnStr
// }

// FOR PRODUCTION BUILD ONLY!!!!

// Init initializes the database using .env vars
func Init() string {
	gotenv.Load()
	url := os.Getenv("DATABASE_URL")
	ConnStr := url
	return ConnStr
}
