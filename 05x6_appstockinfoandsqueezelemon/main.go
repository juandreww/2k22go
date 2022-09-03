package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
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
	fmt.Println("oke")
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
	case rMethod == "GET":
		handleGet(conn)
	case rMethod == "POST":
		handlePost(conn)
	default:
		handleGet(conn)
	}
}

func handleGet(conn net.Conn) {

}

func handlePost(conn net.Conn) {
	
}