// Resources
// https://thenewstack.io/make-a-restful-json-api-go/
// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

package main

import "os"

var a App

func main() {
    a = App{}
    a.Initialize(
        os.Getenv("GAPI_DB_USERNAME"),
        os.Getenv("GAPI_DB_PASSWORD"),
        os.Getenv("GAPI_DB_NAME"),
    )

    a.Run(":8080")
}

func GiveMeCurrentApp() (a App) {
    return a
}