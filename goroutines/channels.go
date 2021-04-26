package main

import (
	"log"
	"math/rand"
)

func main() {
	channel := make(chan int)

	// can be in either order
	go func(chan int) {

		for {
			rand := rand.Intn(200)
			channel <- rand
		}
	}(channel)

	for {
		select {
		case integer := <-channel:
			log.Printf("Received value at end of channel -> %d", integer)
		}
	}
}
