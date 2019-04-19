package main

import (
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

var once sync.Once
var once001 sync.Once
var once002 sync.Once

func main() {
	//onceA()
	/*
		2019/04/19 18:10:32 multiSyncOnce.go:24: onces
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 0
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 1
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 2
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 3
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 4
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 5
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 6
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 7
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 8
		2019/04/19 18:10:32 multiSyncOnce.go:34: count:  --- 9
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
		2019/04/19 18:10:32 multiSyncOnce.go:40: 213
	*/

	onceB()
	/*

	 */
}

func onces() {
	log.Println("onces")
}
func onced() {
	log.Println("onced")
}

func onceB() {
	for i, _ := range make([]string, 10) {
		once001.Do(onces)
		log.Println("count:", "---", i)
	}
	for i := 0; i < 10; i++ {

		go func(i int) {
			once002.Do(onced)
			log.Println("count:", "---", i)
		}(i)
		time.Sleep(100)
	}
	time.Sleep(4000)
}

func onceA() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		log.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			log.Println("213")
		}()
	}
	time.Sleep(4000)
}
