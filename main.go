package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app, err := newDefaultApp()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	app.registerRoute(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
