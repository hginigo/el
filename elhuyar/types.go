package elhuyar

type WordResult struct {
    Words           []string        `json:"words"`
    Translations    []Translation   `json:"translations"`
}

type Translation struct {
    From    string  `json:"from"`
    To      string  `json:"to"`
    Entries []Entry `json:"entries"`
}

type Entry struct {
    Sort        string      `json:"sort"`
    Entry       []string    `json:"entry"`
    Examples    []Example   `json:"examples"`
}

type Example struct {
    Sentence     string  `json:"sentence"`
    Translation  string  `json:"translation"`
}

const (
    Red = "\033[31m"
    Yel = "\033[33m"
    Blu = "\033[34m"
    Whi = "\033[37m"
    Def = "\033[39;49m"
    Void = ""
)

