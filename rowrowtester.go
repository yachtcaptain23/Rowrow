package main

import (
  "net"
  "fmt"
  "log"
  "time"
)

func printToTcp(urlport string) {
    // Listen on TCP port 1234
    fmt.Println("Writing to %s", urlport)
    addr, _ := net.ResolveTCPAddr("tcp", urlport)
    conn, err := net.DialTCP("tcp", nil, addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    err = conn.SetNoDelay(true)
    if err != nil {
        log.Fatal(err)
    }
    // Wait for a connection.
    var message string
    message = "albert wang"
    conn.Write([]byte(message))
    fmt.Println("Wrote to pipeline")
}

func main() {
  go printToTcp("localhost:1234")
  time.Sleep(30 * time.Second)
}
