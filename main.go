package main

import (
  "net/http"
  "log"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
)

func main() {

  setupStorage()

	http.HandleFunc("/registerDevice", registerDevice)
  http.HandleFunc("/subscribe", subscribe)
  http.HandleFunc("/unsubscribe", unsubscribe)

	logError(http.ListenAndServe(":8080", nil))
}

func setupStorage() {

  deviceStatement, _ := database().Prepare("CREATE TABLE IF NOT EXISTS devices (id INTEGER PRIMARY KEY, token TEXT)")

	_, error := deviceStatement.Exec()
	logError(error)

  if error == nil {

      subscriptionsStatement, _ := database().Prepare("CREATE TABLE IF NOT EXISTS subscriptions (id INTEGER PRIMARY KEY, symbol TEXT, device_id INTEGER, FOREIGN KEY(device_id) REFERENCES devices(id))")
      _, error := subscriptionsStatement.Exec()
      logError(error)
  }
}

func database() *sql.DB {

	database, error := sql.Open("sqlite3", "./storage.db")
	logError(error)

	return database
}

func registerDevice(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "register device")
}

func subscribe(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "subscribe")
}

func unsubscribe(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "unsubscribe")
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
