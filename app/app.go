package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.ensureTableExists()

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/notification/{id:[0-9]+}", a.getNotification).Methods("GET")
	a.Router.HandleFunc("/notification", a.createNotification).Methods("POST")

}

func (a *App) ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS notifications
(
 id SERIAL,
 number NUMERIC,
 message TEXT,
 CONSTRAINT notifications_pkey PRIMARY KEY (id)
)`
