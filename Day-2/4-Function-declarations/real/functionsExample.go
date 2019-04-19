package main

import (
	"log"
	"reflect"
)

type Handler func() *int
type VarHandler func(...int)
type IntHandler func(int) int

// function compositions
func compose(a IntHandler, b IntHandler) IntHandler {
	return func(c int) int {
		return a(b(c))
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func init() {
	log.Println("Call-1")
}

func init() {
	log.Println("Call-2")
}

func main() {
	var f Handler = func() *int {
		i := 1
		return &i
	}

	//log.Println(&f()) //cannot take a ref of function call
	log.Println(&f)

	log.Println(reflect.TypeOf(f).Comparable()) // false

	var vh VarHandler = func(i ...int) {
		log.Println(i)
	}

	vh([]int{1, 2, 3}...)

	add2 := compose(func(i int) int {
		return i + 1
	}, func(i int) int {
		return i + 1
	})

	log.Println(add2(0)) // 2
}
