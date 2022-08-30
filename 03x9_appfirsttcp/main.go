package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"net/http"
	"encoding/json"
)

type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

func main() {
	// mux := http.DefaultServeMux
	// mux.HandleFunc("/first", first())
	// mux.HandleFunc("/second", second())
	// http.ListenAndServe(":8080", mux)

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
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
	// respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			mux(conn, ln)
		}
		if ln == "" {
			break
		}

		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	if m == "GET" && u == "/first" {
		first(conn)
	}

	if m == "GET" && u == "/second" {
		second(conn)
	}
}

func first(conn net.Conn) {

}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}