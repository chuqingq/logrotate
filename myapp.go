package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Printf("hello world")
		time.Sleep(time.Millisecond)
	}
}
