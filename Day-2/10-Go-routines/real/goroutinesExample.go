package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	log.Println(runtime.NumCPU()) // logical CPUs
	go func() {
		select {} // block
	}()

	log.Println(runtime.NumGoroutine()) // 2 = main + inf select

	ir := func() int {
		var i = 1
		log.Println(i)
		return i
	}

	go func(n int) {
		log.Println("Got,", n)
	}(ir())

	go func(n int) {
		log.Println("Got,", n)
	}(ir())

	go func(n int) {
		log.Println("Got,", n)
	}(ir())

	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	var v int32 = 0

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1) // wrong place, must be called before firing goroutine
			atomic.AddInt32(&v, 1)
			wg.Done() // or wg.Add(-1)
		}()
	}
	wg.Wait()
	log.Println(atomic.LoadInt32(&v)) // might print < 100
}
