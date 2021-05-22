package main

import (
	"fmt"
	"os"
	"encoding/json"
        . "github.com/hginigo/el/elhuyar"
)

func main() {

    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Fprintln(os.Stderr, "Errorea: ez da argumenturik eman")
        return
    }

    res, err := FetchResult(args[0])
    if err != nil {
        fmt.Fprint(os.Stderr, "Errorea")
        return
    }

    enc := json.NewEncoder(os.Stdout)
    enc.SetIndent("", "    ")
    enc.Encode(res)

    // res, err := json.Marshal(queryList)
    // if err != nil {
    //     fmt.Fprintln(os.Stderr, "Unexpected error")
    //     return
    // }
}

