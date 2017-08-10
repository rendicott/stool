package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
    "strconv"
    "io"
    "io/ioutil"
)

type App struct {
    Router  *mux.Router
    DB      *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

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

func GameIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(GetGames(a.DB)); err != nil {
        panic(err)
    }
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gameId, err := strconv.Atoi(vars["gameId"])
    if err != nil {
        panic(err)
    }

    g := Game{Id: gameId}
    if errr := g.GetGame(a.DB); errr != nil {
        panic(errr)
    }
    if errrr := json.NewEncoder(w).Encode(g); errrr != nil {
        panic(errrr)
    }
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
    var g Game
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if errr := r.Body.Close(); errr != nil {
        panic(errr)
    }
    if errrr := json.Unmarshal(body, &g); errrr != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if errrrr := json.NewEncoder(w).Encode(errrr); errrrr != nil {
            panic(errrrr)
        }
    }
    if errrrrr := g.CreateGame(a.DB); errrrrr != nil {
        panic(errrrrr)
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if errrrrrr := json.NewEncoder(w).Encode(g); errrrrrr != nil {
        panic(errrrrrr)
    }
}