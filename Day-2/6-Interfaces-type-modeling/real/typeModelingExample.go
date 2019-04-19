package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type Cleaner interface {
	Clean() bool
}

type Eraser interface {
	Erase() bool
}

// Type Embedding
type Destroyer interface {
	Cleaner
	Eraser
}

// Pointer to pointer to string Left -> Right
type PPS = **string
type WebController struct{}

func (wc *WebController) GetName() string {
	return "Web Controller"
}

type Indexer interface {
	Index()
}

// Anonymous type embedding
type AppController struct {
	*WebController
	Indexer

	// Will not compile
	// PPS
	// *PPS
	// *Indexer
}

type IndexString string

func (hs IndexString) Index() {
	if len(hs) != 0 {
		log.Println(hs)
	} else {
		log.Println("Index Page")
	}

}

func main() {
	ac := new(AppController)
	log.Println(ac.WebController.GetName())
	// shorthand
	log.Println(ac.GetName())
	//panics
	//ac.Index()
	ac = &AppController{new(WebController), IndexString("No. 100")}
	ac.Index()
}
