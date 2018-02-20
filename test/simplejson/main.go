package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	str := new(Sabdakosh)
	err := json.Unmarshal([]byte(input), &str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
	fmt.Printf(str.Word)
}
