package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
	"io"
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

		go serve(conn)
	}
	// fmt.Println("oke")
}

func serve(conn net.Conn) {
	defer conn.Close()
	var i int
	var rMethod, rURL string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURL = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URL:", rURL)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	switch {
	case rMethod == "GET" && rURL == "/openicebox":
		handleGet(conn)
	case rMethod == "POST" && rURL == "/squeezelemon":
		handlePost(conn)
	default:
		handleDefault(conn)
	}
}

func handleGet(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>CHECK LEMON STOCK</title>
		</head>
		<body>
			<h1>"CHECK LEMON STOCK"</h1>
			<a href="/">Home</a><br>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handlePost(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>SQUEEZED SOME LEMONS</title>
		</head>
		<body>
			<h1>"SQUEEZED SOME LEMONS"</h1>
			<a href="/">Home</a><br>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefault(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>WELCOME TO LEMONADE STAND 1819</title>
		</head>
		<body>
			<h1>"HI ANDREW, WHAT ACTIVITY YOU WANT TO DO?"</h1>
			<a href="/openicebox">Open Ice Box</a><br>
			<a href="/squeezelemon">Squeeze some lemons</a><br>
			<form action="/squeezelemon" method="POST">
			<input type="hidden" value="squeeze the lemon please">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}