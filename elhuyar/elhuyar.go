package elhuyar

import (
    "io"
    "os"
    "fmt"
    "time"
    "sync"
    "strings"
    "github.com/gocolly/colly"
)

const ELH_URL = "https://hiztegiak.elhuyar.eus/eu_es_en_fr/"

func FetchResult(query string) ([]Translation, error) {
    translationList := make([]Translation, 0)
    c := colly.NewCollector(
        colly.AllowedDomains("hiztegiak.elhuyar.eus", "www.hiztegiak.elhuyar.eus"),
    )

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
                ex_vec = append(ex_vec, Example{ Sentence: sentences[0], Translation: sentences[1] })
            })

            entry := Entry {
                Sort: sort,
                Entry: entries,
                Examples: ex_vec,
            }

            entryList = append(entryList, entry)
        })
        translation := Translation {
            From: from,
            To: to,
            Entries: entryList,
        }
        translationList = append(translationList, translation)
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

    c.Visit(ELH_URL + query)
    wg.Wait()

    return translationList, nil
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

func FprintResult(ts []Translation, w io.Writer) {
    var col string
    var t Translation
    var e Entry
    var x Example

    for i := range ts {
        t = ts[i]
        col = getColor(t.From, t.To)
        fmt.Fprintf(w, "%s| %s > %s%s\n", col, t.From, t.To, Def)

        for j := range t.Entries {
            e = t.Entries[j]
            fmt.Fprintf(w, "%s|  %d %s%s %s%s\n", col, j+1, e.Sort, Whi, fmtStrList(e.Entry), Def)

            for k := range e.Examples {
                x = e.Examples[k]
                fmt.Fprintf(w, "%s|    %s%s%s -> %s%s\n", col, Whi, x.Sentence, col, Def, x.Translation)
            }
            // Do not print separation line if it is the last entry
            if j < len(t.Entries) - 1 {
                fmt.Fprintf(w, "%s|%s\n", col, Def)
            }
        }
        fmt.Fprintln(w)
    }
}

func fmtStrList(s []string) string {
    if len(s) == 0 {
        return ""
    }
    var res string = s[0]
    for i := range s[1:] {
        res += ", " + s[i]
    }
    return res
}

func getColor(from, to string) string {
    if from == "es" || to == "es" {
        return Red
    } else if from == "fr" || to == "fr" {
        return Blu
    }
    return Yel
}
