package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// lets define our mgo server first
	srv, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	// creates a new top level mux.Router. since a mux.Router implements the http.Handler interface,
	// we can pass it to http.ListenAndServe below
	router := mux.NewRouter()

	// configure the router to always run this handler when it couldn't match a request to any other handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s not found\n", r.URL)))
	})

	// create a subrouter just for standard API calls. subrouters are convenient ways to
	// group similar functionality together. this subrouter also verifies that the Content-Type
	// header is correct for a JSON API.
	apiRouter := router.Headers("Content-Type", "application/json").Subrouter()
	apiRouter.HandleFunc("/api/sabda/{word}", srv.WithData(getSabda)).Methods("GET")
	apiRouter.HandleFunc("/api/sabdas", srv.WithData(getAllSabda)).Methods("GET")
	log.Printf("serving on port http://localhost:2048")
	http.ListenAndServe("localhost:2048", router)

}

// func ensureIndex(s *session) goji.HandlerFunc {
// 	return func(ctx context.Cntext, w http.ResponseWriter, r *http.Request) {
// 		session := s.Copy()
// 		defer session.Close()
// c := session.DB(os.Getenv("DBNAME")).C(os.Getenv("COLNAME"))

// 	}
// }
