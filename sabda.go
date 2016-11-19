package main

type Noun struct {
	Meaning []string `bson:"m" json:"m,omitempty"`
	Tags    []string `bson:"_#" json:"_#,omitempty"`
}
type Verb struct {
	Meaning []string `bson:"m" json:"m,omitempty"`
	Tags    []string `bson:"_#" json:"_#,omitempty"`
}

type POS struct {
	Noun []Noun `bson:"n" json:"n,omitempty"`
	Verb []Verb `bson:"v" json:"v,omitempty"`
}

type Metas struct {
	Sci   string   `bson:"_sci" json:"_sci,omitempty"`
	Wikis []string `bson:"_wiki" json:"_wiki,omitempty"`
	Tags  []string `bson:"_#" json:"_#,omitempty"`
}

type Sabdakosh struct {
	Word  string `bson:"w" json:"w,omitempty"`
	POS   `bson:",inline" json:",inline"`
	Trans []string `bson:"_t" json:"_t,omitempty"`
	Metas `bson:",inline" json:",inline"`
}
