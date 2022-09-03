package main

import (
	"fmt"
	// "net/http"
	"net"
	"io"
	"log"
	"bufio"
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
		go handle(conn)

		
		io.WriteString(conn, "Hello from TCP server\n")
		fmt.Fprintln(conn, "Okay youre done")

		conn.Close()
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("code got here")
}