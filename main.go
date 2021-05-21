package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"encoding/json"
	"github.com/gocolly/colly"
)

type Query struct {
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
    Example     string  `json:"example"`
    Traduction  string  `json:"traduction"`
}

func main() {

    c:= colly.NewCollector(
        colly.AllowedDomains("hiztegiak.elhuyar.eus", "www.hiztegiak.elhuyar.eus"),
    )

    queryList := make([]Query, 0)
    c.OnHTML("div.box_def:has(div.innerDef)", func(e *colly.HTMLElement) {
        names := e.ChildText("h2")
        lang_vec := strings.Split(names, " > ")
        from, to := lang_vec[0], lang_vec[1]
        entryList := make([]Entry, 0)

        e.ForEach("li", func(i int, e *colly.HTMLElement) {
            sort := e.ChildText("p.lehena em[title]")
            entries := make([]string, 0)

            e.ForEach("p.lehena a, span:not(.fina)", func(i int, e *colly.HTMLElement) {
                entries = append(entries, e.Text)
            })

            var ex_vec []Example = make([]Example, 0)

            e.ForEach("div.padDefn > p", func(i int, e *colly.HTMLElement) {
                sentences := strings.Split(e.Text, ": ")
                ex_vec = append(ex_vec, Example{ sentences[0], sentences[1] })
            })

            entry := Entry {
                Sort: sort,
                Entry: entries,
                Examples: ex_vec,
            }

            entryList = append(entryList, entry)
        })
        query := Query {
            From: from,
            To: to,
            Entries: entryList,
        }
        queryList = append(queryList, query)
    })

    var wg sync.WaitGroup
    wg.Add(1)
    c.OnRequest(func(r *colly.Request) {
        go printWait()
    })

    c.OnResponse(func(r *colly.Response) {
        wg.Done()
        fmt.Printf("\033[1K\n\033[A")
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Fprintln(os.Stderr, "URL eskakizuna:", r.Request.URL, "erantzuna:", r, "\nError:", err)
    })

    c.Visit("https://hiztegiak.elhuyar.eus/eu_es/esan")
    wg.Wait()

    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", " ")
    enc.Encode(queryList)

    // res, err := json.Marshal(queryList)
    // if err != nil {
    //     fmt.Fprintln(os.Stderr, "Unexpected error")
    //     return
    // }
}

func printWait() {
    symbols := [3]string{ "/", "-", "\\" }
    i := 0;

    fmt.Print("Bilatzen... ")
    for {
        fmt.Printf("\033[D%s", symbols[i])
        i = (i + 1) % len(symbols)
        time.Sleep(time.Millisecond * 100)
    }
}
