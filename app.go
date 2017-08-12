package main

//to add: curl -H "Content-Type: application/json" -d '{"name": "Camel Up"}' http://localhost:8080/games
//to delete: curl -X "DELETE" http://localhost:8080/players/4

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

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GameIndex",
        "GET",
        "/games",
        GameIndex,
    },
    Route{
        "ShowGame",
        "GET",
        "/games/{gameId}",
        ShowGame,
    },
    Route{
        "GameCreate",
        "POST",
        "/games",
        CreateGame,
    },
    Route{
        "PlayerIndex",
        "GET",
        "/players",
        PlayerIndex,
    },
    Route{
        "ShowPlayer",
        "GET",
        "/players/{playerId}",
        ShowPlayer,
    },
    Route{
        "PlayerCreate",
        "POST",
        "/players",
        CreatePlayer,
    },
    Route{
        "GameDelete",
        "DELETE",
        "/games/{gameId}",
        a.DeleteGame,
    },
    Route{
        "PlayerDelete",
        "DELETE",
        "/players/{playerId}",
        a.DeletePlayer,
    },
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to the gAPI")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(GetGames(a.DB)); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gameId, err := strconv.Atoi(vars["gameId"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid ID")
        return
    }

    g := Game{Id: gameId}
    if errr := g.GetGame(a.DB); errr != nil {
        respondWithError(w, http.StatusInternalServerError, errr.Error())
        return
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
        respondWithError(w, http.StatusInternalServerError, errrrrr.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if errrrrrr := json.NewEncoder(w).Encode(g); errrrrrr != nil {
        //respond with error here
        panic(errrrrrr)
    }
}

func (a *App) DeleteGame(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gameId, err := strconv.Atoi(vars["gameId"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid ID")
        return
    }

    g := Game{Id: gameId}
    if err := g.DeleteGame(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
}

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(GetPlayers(a.DB)); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
}

func ShowPlayer(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    playerId, err := strconv.Atoi(vars["playerId"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid ID")
        return
    }

    p := Player{Id: playerId}
    if errr := p.GetPlayer(a.DB); errr != nil {
        respondWithError(w, http.StatusInternalServerError, errr.Error())
        return
    }
    if errrr := json.NewEncoder(w).Encode(p); errrr != nil {
        panic(errrr)
    }
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
    var p Player
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if errr := r.Body.Close(); errr != nil {
        panic(errr)
    }
    if errrr := json.Unmarshal(body, &p); errrr != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if errrrr := json.NewEncoder(w).Encode(errrr); errrrr != nil {
            panic(errrrr)
        }
    }
    if errrrrr := p.CreatePlayer(a.DB); errrrrr != nil {
        respondWithError(w, http.StatusInternalServerError, errrrrr.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if errrrrrr := json.NewEncoder(w).Encode(p); errrrrrr != nil {
        respondWithError(w, http.StatusInternalServerError, errrrrrr.Error())
        return
    }
}

func (a *App) DeletePlayer(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    playerId, err := strconv.Atoi(vars["playerId"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid ID")
        return
    }

    p := Player{Id: playerId}
    if err := p.DeletePlayer(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
}