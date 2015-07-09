package main

import (
	. "github.com/zubairhamed/minibus"
	"log"
	"math/rand"
	"time"
)

func main() {
	bus := NewMiniBus()

	bus.Sub("topic", func(topic interface{}) {
		log.Println("Calling from inside subscription: ", topic)
	})

	rand.Seed(50)
	for {
		time.Sleep(1 * time.Second)

		v := rand.Intn(2)
		if v == 1 {
			bus.Pub("topic", "This is a topic message.")
		}
	}
}
