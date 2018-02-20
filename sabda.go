package main

type Noun struct {
	Meaning         []string `bson:"tr,omitempty" json:"tr,omitempty"`
	Tags            []string `bson:"#,omitempty" json:"#,omitempty"`
	Transliteration []string `bson:"t,omitempty" json:"t,omitempty"`
}
type Verb struct {
	Meaning         []string `bson:"tr,omitempty" json:"tr,omitempty"`
	Tags            []string `bson:"#,omitempty" json:"#,omitempty"`
	Transliteration []string `bson:"t,omitempty" json:"t,omitempty"`
}
type Adjective struct {
	Meaning         []string `bson:"tr,omitempty" json:"tr,omitempty"`
	Tags            []string `bson:"#,omitempty" json:"#,omitempty"`
	Transliteration []string `bson:"t,omitempty" json:"t,omitempty"`
}

type POS struct {
	Noun      []Noun      `bson:"n,omitempty" json:"n,omitempty"`
	Adjective []Adjective `bson:"j,omitempty" json:"j,omitempty"`
	Verb      []Verb      `bson:"v,omitempty" json:"v,omitempty"`
}

type Metas struct {
	Sci   string   `bson:"sci,omitempty" json:"sci,omitempty"`
	Wikis []string `bson:"_wiki,omitempty" json:"_wiki,omitempty"`
	Tags  []string `bson:"#,omitempty" json:"#,omitempty"`
}

type Sabdakosh struct {
	Word  string `bson:"w" json:"w,omitempty"`
	POS   `bson:",inline" json:",inline"`
	Metas `bson:",inline" json:",inline"`
}
