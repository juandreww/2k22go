package main

import (
	"fmt"
	// "net/http"
	"net"
	"io"
	"log"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer li.Close()
		io.WriteString(conn, "Hello from TCP server\n")
		fmt.Fprintln(conn, "Okay youre done")

		conn.Close()
	}

}