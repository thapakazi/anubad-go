package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Anubad struct {
	Word    string
	Meaning string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		anubad := Anubad{Word: "Dictonary", Meaning: "शब्दकोष"}
		if word := r.FormValue("sabda"); word != "" {
			anubad.Word = word
		}
		if err := templates.ExecuteTemplate(w, "index.html", anubad); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	fmt.Println(http.ListenAndServe(":2048", nil))
}
