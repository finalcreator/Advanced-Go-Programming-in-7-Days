package main

import (
	"log"
	"runtime/debug"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// main go-routine
	log.Println("\n" + string(debug.Stack()))
	defer debug.PrintStack()
	done := make(chan bool)

	// from goroutine
	log.Print("\n\nfrom goroutine\n\n")
	go func(done chan bool) {
		debug.PrintStack()
		done <- true
		close(done)
	}(done)
	<-done

	// using Stack()
	log.Print("\n\nusing Stack()\n\n")
	stackTrace := debug.Stack()
	log.Printf("%v", string(stackTrace))

}
