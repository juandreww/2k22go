package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"net/http"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	mux := http.DefaultServeMux
	mux.HandleFunc("/first", first())
	mux.HandleFunc("/second", second())

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func first() string {
	return "first"
}

func second() string {
	return "second"
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// fmt.Println(i)
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			c := strings.Fields(ln)[1]
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", c)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}