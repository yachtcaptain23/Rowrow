package main

import (
  "net"
  "io"
  "fmt"
  "log"
)

func tcpHandling(urlport string) {
    // Listen on TCP port 1234
    fmt.Println("Listening to %s", urlport)
    tcpSock, err := net.Listen("tcp", urlport)
	if err != nil {
		log.Fatal(err)
	}
	defer tcpSock.Close()
	for {
		// Wait for a connection.
		conn, err := tcpSock.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

func whatisgoingon(haha string) {
  fmt.Println("this is going on=%s", haha)
}

func unixHandling(fileptr string) {
    // Read from unix domain socket
    fmt.Println("Listening to %s", fileptr)
    unixSock, err := net.Listen("unix", fileptr)
    if err != nil {
    log.Fatal(err)
    }
    defer unixSock.Close()
    for {
        // Wait for a connection.
        conn, err := unixSock.Accept()
        if err != nil {
            log.Fatal(err)
        }
        // Handle the connection in a new goroutine.
        // The loop then returns to accepting, so that
        // multiple connections may be served concurrently.
        go func(c net.Conn) {
            // Echo all incoming data.
            io.Copy(c, c)
            // Shut down the connection.
            c.Close()
        } (conn)
    }
}

func main() {
    whatisgoingon("programming")
    go tcpHandling("localhost:1234")
    go unixHandling("/tmp/ipc_sock")
    whatisgoingon("trolol")

    // Everytime the Unix socket sends a new packet, close the old 
    // TCP socket, re-open a new one and re-send the data to localhost:1234
}
