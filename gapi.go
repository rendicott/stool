package main

// use go get first
// Also had to export GOBIN=$GOPATH/bin
import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/games", ShowGames)
    router.HandleFunc("/games/{testId}", Test)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome to the gAPI")
}

func ShowGames(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "SHOW ME THE GAMES")
}

func Test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testId := vars["testId"]
	fmt.Fprintln(w, "Test show:", testId)
}