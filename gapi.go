package main

// use go get first
// Also had to export GOBIN=$GOPATH/bin
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/games", GameIndex)
    router.HandleFunc("/games/{testId}", Test)
    router.HandleFunc("/gameindex", GameIndex)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to the gAPI")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
	games := Games{
		Game{Name: "Splendor", Players: 4, Id: 1},
		Game{Name: "Love Letter", Players: 4, Id: 2},
	}

	json.NewEncoder(w).Encode(games)
}

func Test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testId := vars["testId"]
	fmt.Fprintln(w, "Test show:", testId)
}

type Game struct {
	Name 		string
	Players		int
	Id 			int
}

type Games []Game
