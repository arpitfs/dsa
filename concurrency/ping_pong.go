package concurrency

import (
	"fmt"
	"time"
)

func pingPong() {
	pingChannel := make(chan bool)
	pongChannel := make(chan bool)

	go printPing(pingChannel, pongChannel)
	go printPong(pingChannel, pongChannel)

	pingChannel <- true
	time.Sleep(5 * time.Second)

}

func printPing(ping, pong chan bool) {
	for {
		<-ping
		fmt.Println("Ping")
		time.Sleep(500 * time.Millisecond)
		pong <- true
	}
}

func printPong(ping, pong chan bool) {
	for {
		<-pong
		fmt.Println("Pong")
		time.Sleep(500 * time.Millisecond)
		ping <- true
	}
}
