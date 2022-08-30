package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	fmt.Println("Phase 1")
	for {
		// LAC : Listen Accept Connect
		fmt.Println("Phase 2")
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Phase 3")
		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well I Hope")

		fmt.Println("Phase 4")
		conn.Close()
	}
}