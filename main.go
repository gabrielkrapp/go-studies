package main

import (
	"database/sql"
	dbconfig "estudos-go/model/config"
	"estudos-go/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var err error
	a.DB, err = dbconfig.DatabaseConnect()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/register", routes.Register(a.DB)).Methods("POST")
	a.Router.HandleFunc("/editpassword", routes.EditPasswordRequest(a.DB)).Methods("POST")
	a.Router.HandleFunc("/deleteuser", routes.DeleteUserRequest(a.DB)).Methods("DELETE")
	a.Router.HandleFunc("/getusers", routes.GetUsersRequest(a.DB)).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	app := &App{}
	app.Initialize()
	defer app.DB.Close()
	app.Run(":8080")
}
