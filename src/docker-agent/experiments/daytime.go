package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError1(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError1(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        daytime := time.Now().String()
        fmt.Println(daytime)

        conn.Write([]byte(daytime)) // don't care about return value
        conn.Close()                // we're finished with this client
    }
}

func checkError1(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}