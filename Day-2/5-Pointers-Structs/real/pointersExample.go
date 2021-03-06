package main

import (
	"log"
)

type P int

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func add2(a int) func() int {
	return func() int {
		return a + 2
	}
}

func add2Ref(a *int) func() int {
	return func() int {
		return *a + 2
	}
}

func add2Ref2(a *int) func() int {
	var b = *a // save inside closure
	return func() int {
		return b + 2
	}
}

func main() {
	n := new(P)
	*n = 1
	log.Println(&*n == n)
	//log.Println(*n++) // no direct pointer arithmetic (use unsafe.Pointer)

	a2 := add2(1)
	log.Println(a2()) // 3

	var a0 = 1
	closure0 := add2(a0)
	a0 = 6
	result0 := closure0()
	log.Println(result0)

	var a = 6
	a2r := add2Ref(&a)
	a = 9
	log.Println(a2r()) //11

	a = 3
	a2r2 := add2Ref2(&a)
	a = 9
	log.Println(a2r2()) //5

	s1 := [...]int{1, 2, 3, 4, 5}
	// s is pointer to array
	for _, v := range &s1 {
		log.Println(v)
	}

	//s2 := []int{1,2,3,4,5}
	// wont work on slices
	//for _, v := range &s2 {
	//	log.Println(v)
	//}
}
