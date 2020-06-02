package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    resp, err := http.Get("https://arwildo.github.io")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Response Code: ", resp.StatusCode, http.StatusText(resp.StatusCode))
}
