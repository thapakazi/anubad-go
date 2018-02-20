package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

func homeHandler(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Yay! We're home, Jim!"))
}

func getAllTheThings(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Look, Jim! Get all the things!"))
}

func putOneThing(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Look, Jim! Put one thing!"))
}

func deleteOneThing(wr http.ResponseWriter, req *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Look, Jim! Delete one thing!"))
}

func main() {
	router := pat.New()

	router.Get("/things", getAllTheThings)

	router.Put("/things/{id}", putOneThing)
	router.Delete("/things/{id}", deleteOneThing)

	router.Get("/", homeHandler)

	http.Handle("/", router)

	log.Print("Listening on 127.0.0.1:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
