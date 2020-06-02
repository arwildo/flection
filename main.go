package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "io/ioutil"
    "strings"
)

var (
    Black   = "\033[1;30m%s\033[0m"
    Red     = "\033[1;31m%s\033[0m"
    Green   = "\033[1;32m%s\033[0m"
    Yellow  = "\033[1;33m%s\033[0m"
    Purple  = "\033[1;34m%s\033[0m"
    Magenta = "\033[1;35m%s\033[0m"
    Teal    = "\033[1;36m%s\033[0m"
    White   = "\033[1;37m%s\033[0m"
)

func main() {
    arg := os.Args
    if len(arg) > 1 {
        resp, err := http.Get(arg[1])
        if err != nil {
            log.Fatal(err)
        }

        // Get Respon
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        // Search for keyword
        if strings.Contains(string(body), "FUZZ") {
            fmt.Println(arg[1], "  ", resp.StatusCode, http.StatusText(resp.StatusCode), "  ", "\033[1;32mFound Reflection!\033[0m")
        } else {
            fmt.Println(arg[1], "  ", resp.StatusCode, http.StatusText(resp.StatusCode), "  ", "\033[1;33mNo Reflection\033[0m")
        }

    } else {
        fmt.Println("No arguments were specified. See help -h.")
    }
}
