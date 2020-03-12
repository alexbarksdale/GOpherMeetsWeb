package main

import (
	"bufio"
	"fmt"
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
		go handle(conn) // send connection to this func
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	// Never gets here
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// How does the reader know when it's done?
	fmt.Println("We got here.")
}
