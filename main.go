package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":7000", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	// saveTimeToDatabase()

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to SQLite database!")
}