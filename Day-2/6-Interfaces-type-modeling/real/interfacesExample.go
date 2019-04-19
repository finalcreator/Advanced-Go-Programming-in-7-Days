package main

import (
	"log"
	"reflect"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// type alias
type Helper = interface {
	Help() string
}

type HelpString string

func (hs HelpString) Help() string {
	return string(hs)
}

type UnHelpString struct{}

func (uhs *UnHelpString) Help() string {
	return "I cannot help you"
}

// Compile time check
var x = Helper(HelpString("Not Any"))

func main() {
	log.Println(x)                        //Not Any
	log.Println(reflect.TypeOf(x))        //main.HelpString
	log.Println(HelpString("Hey").Help()) //Hey
	log.Println((&UnHelpString{}).Help()) //I cannot help you
	for _, helper := range []Helper{HelpString("Tom"), &UnHelpString{}} {
		log.Println(helper.Help())
	}
	//Tom
	//I cannot help you

	var h Helper = HelpString("Help me")
	log.Print(h.Help())

	var exp = interface{ Help() string }.Help(h)
	log.Println(exp)

	//Polymorphism
	var helpers = []Helper{
		HelpString("A1"),
		&UnHelpString{},
	}
	log.Println(helpers)

	for _, h := range helpers {
		log.Println(h.Help())
	}

	var h2 interface{} = HelpString("How do you do")
	x, y := h2.(Helper)
	z := h2.(Helper)
	log.Printf("\n%v\n%v\n%v\n", x, y, z)

	//var _ = h2.(string) //panic

}
