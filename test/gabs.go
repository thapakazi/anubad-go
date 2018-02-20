package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
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

func main() {
	jsonParsed, _ := gabs.ParseJSON([]byte(input))

	fmt.Println(jsonParsed.Path("w").Data().(string))
	children, _ := jsonParsed.S("_wiki").Children()
	for _, child := range children {
		fmt.Println(child.Data().(string))
	}
}
