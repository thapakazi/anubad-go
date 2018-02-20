package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Get all sabda
// GET /api/sabdas
func getAllSabda(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "sessionCopy").(*mgo.Session)

	c := session.DB(os.Getenv("DBNAME")).C(os.Getenv("COLNAME"))
	var sabdas []Sabdakosh

	err := c.Find(bson.M{}).All(&sabdas)
	// 		fmt.Printf("%+v", sabdas)
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all sabdas: ", err)
		return
	}

	respBody, err := json.MarshalIndent(sabdas, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)

}

// get single workd(sabda) like
// GET /api/sabda/{word}
func getSabda(w http.ResponseWriter, r *http.Request) {

	// lets get the context first
	session := context.Get(r, "sessionCopy").(*mgo.Session)
	c := session.DB(os.Getenv("DBNAME")).C(os.Getenv("COLNAME"))

	// mux.Vars gets a map of path variables by name. here "name" matches the {name} path
	// variable as seen in gorilla_server.go
	word, ok := mux.Vars(r)["word"]
	if !ok {
		http.Error(w, "name missing in URL path", http.StatusBadRequest)
		return
	}

	var sabda Sabdakosh

	err := c.Find(bson.M{"w": word}).One(&sabda)
	fmt.Printf("%+v", sabda)
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed get all sabda: ", err)
		return
	}

	respBody, err := json.MarshalIndent(sabda, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	// params := word + " inside getSabda"
	// respBody := []byte(params)
	ResponseWithJSON(w, respBody, http.StatusOK)
	// }
}

func addSabda(w http.ResponseWriter, r *http.Request) {

	var sabda Sabdakosh
	type MongoKosh struct {
		sabda Sabdakosh     `bson:",inline," json:",inline,omitempty"`
		ID    bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	//first read the body of request with a json decoder
	decoder := json.NewDecoder(r.Body)
	// fmt.Printf("%+v", decoder)
	// fmt.Printf("%+v", r.Body)
	// then we decode what inside into a struct i.e, anubad
	if err := decoder.Decode(&sabda); err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	msabda := MongoKosh{sabda: sabda}
	session := context.Get(r, "sessionCopy").(*mgo.Session)
	c := session.DB(os.Getenv("DBNAME")).C(os.Getenv("COLNAME"))
	if err := c.Insert(sabda); err != nil {
		if mgo.IsDup(err) {
			ErrorWithJSON(w, "duplicate entry, may be record already present in db", http.StatusBadRequest)
			return
		}
		ErrorWithJSON(w, "database error", http.StatusBadRequest)
		log.Println("Failed to insert new entry", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+string(msabda.ID))
	w.WriteHeader(http.StatusCreated)

}
