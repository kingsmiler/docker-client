package main

import (
    "fmt"
    "time"
    "os"
    "log"
    "net/http"
)

func main() {
    interval := os.Getenv("REPORT_INTERVAL")
    fmt.Println(interval)

    ticker := time.NewTicker(time.Second * 5)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    http.HandleFunc("/", nil)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }


}