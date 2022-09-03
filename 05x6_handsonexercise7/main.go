package main

import (
	"fmt"
	"net"
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

		go serve(conn)
	}
	fmt.Println("oke")
}

func serve(conn net.Conn) {
	defer conn.Close()
}