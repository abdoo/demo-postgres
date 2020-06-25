package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var db *sql.DB

const (
	dbhost = "postgres-headless.keda.svc.cluster.local"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "postgres"
	dbname = "demo"
)

func main() {
	initDb()
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/users", usersHandler)

	http.ListenAndServe(":8383", r)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var count int

	row := db.QueryRow(`SELECT count(*) FROM users`)

	err := row.Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(count)

	fmt.Fprintf(w, "User count is %d", count)
}

func initDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}
