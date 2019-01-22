package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	h "github.com/rafischer1/slr_capstone_go/handlers"
)

// db var references sql.DB
var db *sql.DB

// env variable declarations
const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()

	defer db.Close()
	// port := GetPort()
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         dbport,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// subscription routing
	r.HandleFunc("/subscribe", h.GetAllSubs).Methods("GET")
	r.HandleFunc("/subscribe", h.PostSub).Methods("POST", "OPTIONS")
	r.HandleFunc("/subscribe/{phone}", h.DeleteSub).Methods("DELETE", "OPTIONS")

	// data routing
	r.HandleFunc("/data", h.GetAllData).Methods("GET")
	r.HandleFunc("/data", h.PostData).Methods("POST", "OPTIONS")

	// admin routing
	r.HandleFunc("/admin/{password}", h.AdminVerify).Methods("GET")

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("public/")))

	// set router
	log.Println("Listening...$1", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server stopped")
}

func initDb() {
	// grab .env variables using gotenv package
	gotenv.Load()

	// call dbConfig function to set env variables
	config := dbConfig()
	var err error

	// Loaded database info
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	// open and run the db
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// basic config setup and error handling for db env variables
func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name

	return conf
}
