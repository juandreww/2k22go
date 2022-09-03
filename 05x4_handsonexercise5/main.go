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

		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println(conn, "Information of our shop is all")
			break
		}
	}
	
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	// io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	// io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}