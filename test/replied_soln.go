package main

import (
	"encoding/json"
	"fmt"
	"log"
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

type Meaning struct {
	Translations []string
}

type Translation struct {
	Word            string
	Translations    []string
	ScientificNames []string
	WikiNames       []string
	// map from the category to slice of meanings
	Meanings map[string][]*Meaning
}

func (t *Translation) UnmarshalJSON(buf []byte) error {
	type translation struct {
		Word            string              `json:"w"`
		Translations    []string            `json:"_t"`
		ScientificNames []string            `json:"_sci"`
		WikiNames       []string            `json:"_wiki"`
		Meanings        map[string][]string `json:"_#"`
	}
	// fixed fields
	var tt translation
	if err := json.Unmarshal(buf, &tt); err != nil {
		return err
	}
	// possible meanings (extra keys are ignored)
	var rawMeanings map[string]*json.RawMessage
	if err := json.Unmarshal(buf, &rawMeanings); err != nil {
		return err
	}
	// copy fixed fields over
	//
	// this repetition could be avoided but i'm explicitly keeping the
	// code simpler
	t.Word = tt.Word
	t.Translations = tt.Translations
	t.ScientificNames = tt.ScientificNames
	t.WikiNames = tt.WikiNames

	// translate meanings string pointers into proper types
	meaningNames := make(map[string]*Meaning)
	t.Meanings = make(map[string][]*Meaning, len(tt.Meanings))
	for category, stringpointers := range tt.Meanings {
		fmt.Println("stringpointer:", stringpointers, " category:", category) //n1#spice, n2#vegetable
		for _, strp := range stringpointers {
			fmt.Println("strp: " + strp) // n1 and n2
			m, ok := meaningNames[strp]
			if !ok {
				// we haven't mapped this string pointer yet, do it now
				m = &Meaning{}
				raw, ok := rawMeanings[strp]
				if !ok {
					return fmt.Errorf("dangling meaning string pointer: %q", strp)
				}
				if err := json.Unmarshal(*raw, &m.Translations); err != nil {
					return err
				}
				meaningNames[strp] = m
			}
			t.Meanings[category] = append(t.Meanings[category], m)
		}
	}

	return nil
}

func main() {
	var t Translation
	if err := json.Unmarshal([]byte(input), &t); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", t)
	for category, meanings := range t.Meanings {
		fmt.Printf("\tcategory %q\n", category)
		for _, meaning := range meanings {
			fmt.Printf("\t\t%#v\n", meaning)
		}
	}
}
