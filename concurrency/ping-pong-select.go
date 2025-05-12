package concurrency

import (
	"fmt"
	"time"
)

func pingPongSelect() {
	ping := make(chan string)
	pong := make(chan string)
	go processPing(ping, pong)
	go processPong(ping, pong)

	ping <- "Ping"
	time.Sleep(5 * time.Second)

}

func processPing(ping, pong chan string) {
	for {
		<-ping
		fmt.Println(<-ping)
		pong <- "Pong"
	}
}

func processPong(ping, pong chan string) {
	for {
		<-pong
		fmt.Println(<-pong)
		ping <- "Ping"
	}
}
