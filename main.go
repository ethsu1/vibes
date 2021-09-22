package main

import (
    "vibes"
)
func main() {

	server := &Server{
		port: 8080,
		addresss: "localhost",
	}
    // Listen for incoming connections.
    err := server.Bind()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    for {
        // Listen for an incoming connection.
        conn, err := server.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go server.Serve(conn)
    }
}