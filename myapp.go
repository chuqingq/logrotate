package main 

import (
	"log"
	"time"
	"fmt"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func main() {
	for {
		log.Printf("hello world")
		fmt.Printf("")
		time.Sleep(time.Millisecond)
	}
}
