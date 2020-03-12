package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen for a connection
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// conn - get the connection
		conn, err := li.Accept() // Accept connection
		if err != nil {
			log.Println(err)
		}

		// Write to connection
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Good")
		conn.Close()
	}
}
