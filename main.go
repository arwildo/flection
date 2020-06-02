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
            if err == nil {
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    log.Fatal(err)
                }
                defer resp.Body.Close()

                // Search for keyword
                if strings.Contains(string(body), "FUZZ") {
                    fmt.Printf("%s [%d %s] \033[1;32mFound Reflection!\033[0m\n", scanner.Text(), resp.StatusCode, http.StatusText(resp.StatusCode))
                } else {
                    fmt.Printf("%s [%d %s] \033[1;33mNo Reflection\033[0m\n", scanner.Text(), resp.StatusCode, http.StatusText(resp.StatusCode))
                }
            } else {
                fmt.Printf("%s \033[1;31mCan't be Reached\033[0m\n", scanner.Text())
            }
        }
    } else {
        fmt.Println("No file were specified.")
    }
}
