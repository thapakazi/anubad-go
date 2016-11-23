package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

const input = `
{
    "w": "black mustard",
    "n2": [
        "<leaves>तोरीको साग"
    ],
    "_t": [
        "बल्याक मस्टर्ड"
    ],
    "_#": {
        "#spice": [
            "n1"
        ],
        "#vegetable": [
            "n2"
        ]
    },
    "_sci": [
        "Brassica nigra"
    ],
    "n1": [
        "<plant>तोरी"
    ],
    "_wiki": [
        "Brassica_nigra"
    ]
}
`

type Anubad struct {
	Word   string `json:"w"`
	Means  interface{}
	Trans  []string    `json:"_t"`
	Tags   interface{} `json:"_#"`
	Others bson.M      `bson:",inline" json:"others"`
}
type Tags struct {
	Hash interface{}
}
type Means struct {
	N1 []string
	N2 []string
}

func main() {
	var meanings, tags json.RawMessage
	anubad := Anubad{
		Means: &meanings,
		Tags:  &tags,
	}

	if err := json.Unmarshal([]byte(input), &anubad); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(anubad.Word)
	// fmt.Println(anubad)
	fmt.Printf("%+v\n", anubad)
	// var m Means

	// if err := json.Unmarshal(meanings, &m); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(m)
	var t Tags
	if err := json.Unmarshal(tags, &t); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	// switch env.Type {
	// case "sound":
	// 	var s Sound
	// 	if err := json.Unmarshal(msg, &s); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	var desc string = s.Description
	// 	fmt.Println(desc)
	// default:
	// 	log.Fatalf("unknown message type: %q", env.Type)
	// }
}
