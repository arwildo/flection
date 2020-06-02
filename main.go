package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "io/ioutil"
    "strings"
    "bufio"
)

var (
    Green   = "\033[1;32m%s\033[0m"
    Yellow  = "\033[1;33m%s\033[0m"
)

func main() {
    arg := os.Args

    if len(arg) > 1 {

        // Read file
        file, err := os.Open(arg[1])
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)

        // Send Request for each line
        for scanner.Scan() { 
            resp, err := http.Get(scanner.Text())

            // Get Respon
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                log.Fatal(err)
            }

            // Search for keyword
            if strings.Contains(string(body), "FUZZ") {
                fmt.Println(scanner.Text(), "  ", resp.StatusCode, http.StatusText(resp.StatusCode), "  ", "\033[1;32mFound Reflection!\033[0m")
            } else {
                fmt.Println(scanner.Text(), "  ", resp.StatusCode, http.StatusText(resp.StatusCode), "  ", "\033[1;33mNo Reflection\033[0m")
            }
        }
        if err != nil {
            log.Fatal(err)
        }
    } else {
        fmt.Println("No file were specified.")
    }
}
