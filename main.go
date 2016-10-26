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
	Word    string                 `json:"word"`
	Meaning map[string]interface{} `json:"meaning"`
	Sci     string                 `json:"scientifc_name"`
	Tags    []string               `json:"tags"`
	Spell   string                 `json:"spell"`
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
	http.ListenAndServe("localhost:2048", mux)

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

		respBody, err := json.MarshalIndent(anubad, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
