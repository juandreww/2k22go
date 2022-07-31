package main

import (
	"fmt"
	"time"
)

func pinger()

func ponger()

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	for {
		time.Sleep(time.Second * 1)
	}
}