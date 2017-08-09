package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type App struct {
    Router  *mux.Router
    DB      *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)
    fmt.Printf("Connection string: %s", connectionString)
    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = NewRouter()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":8080", a.Router))
}