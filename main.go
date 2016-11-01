package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type Anubad struct {
	ID      bson.ObjectId          `bson:"_id,omitempty" json:"_id"`
	Word    string                 `json:"word"`
	Meaning map[string]interface{} `json:"meaning"`
	Extra   bson.M                 `bson:",inline" json:"extra"`
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func main() {
	// _ := template.Must(template.ParseFiles("templates/books.html"))

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/sabdaharu"), allAnubads(session))
	mux.HandleFuncC(pat.Get("/sabda/:word"), sabdakhoj(session))
	mux.HandleFuncC(pat.Post("/sabdaharu"), addSabda(session))
	http.ListenAndServe("localhost:2048", mux)

}
func addSabda(s *mgo.Session) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var anubad Anubad

		//first read the body of request with a json decoder
		decoder := json.NewDecoder(r.Body)
		// then we decode what inside into a struct instance
		if err := decoder.Decode(&anubad); err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		//else we put data into the db
		c := session.DB("anubad").C("sabda")
		if err := c.Insert(anubad); err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "duplicate entry, may be record already present in db", http.StatusBadRequest)
				return
			}
			ErrorWithJSON(w, "database error", http.StatusBadRequest)
			log.Println("Failed to insert new entry", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+anubad.Word)
		w.WriteHeader(http.StatusCreated)
	}
}

func allAnubads(s *mgo.Session) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		session := s.Copy()
		defer session.Close()

		c := session.DB("anubad").C("sabda")
		var anubads []Anubad

		err := c.Find(bson.M{}).All(&anubads)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all anubads: ", err)
			return
		}

		respBody, err := json.MarshalIndent(anubads, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func sabdakhoj(s *mgo.Session) goji.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		word := pat.Param(ctx, "word")

		c := session.DB("anubad").C("sabda")

		var anubad Anubad
		err := c.Find(bson.M{"word": word}).One(&anubad)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find word: ", err)
			return
		}

		if anubad.Word == "" {
			ErrorWithJSON(w, "Word not found", http.StatusNotFound)
			return
		}
		fmt.Println(anubad)
		respBody, err := json.MarshalIndent(anubad, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

// func ensureIndex(s *session) goji.HandlerFunc {
// 	return func(ctx context.Cntext, w http.ResponseWriter, r *http.Request) {
// 		session := s.Copy()
// 		defer session.Close()
// 		c := session.DB("anubad").C("sabda")

// 	}
// }
