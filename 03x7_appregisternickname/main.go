package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"math/rand"
	"strconv"
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
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		/*
			insert full name
			check first character of two word after space
			if theres only 1 string then use the first two character
			concat with random int(3)
			upper case letter
		*/
		ln := strings.ToUpper(scanner.Text())
		bs := []byte(ln)
		if len(bs) < 3 {
			log.Fatalln("Your name must be 3 characters minimum")
		}
		r := registernickname(bs)

		fmt.Fprintf(conn, "%s - Your Registered Nickname is: %s\n", ln, r)
		fmt.Fprintln(conn, bs)
	}
}

func registernickname(bs []byte) []byte {
	var res = make([]byte, 5)
	for i, v := range bs {
		if i == 0 {
			res[i] = v
		} else if len(res) == 2 {
			break
		} else if v == 32 && len(bs) > i {
			res[i] = bs[i+1]
		}
	}

	if len(res) == 1 {
		res[1] = bs[1]
	}

	intgr := randomInRange(100, 300)
	str := strconv.Itoa(intgr)

	return res
}

func randomInRange(a, b int) int {
	return a + rand.Intn(b-a)
}