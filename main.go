package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, नमस्ते everyone... this is how api.anubad.io begins")

	})
	fmt.Println(http.ListenAndServe(":2048", nil))
}
