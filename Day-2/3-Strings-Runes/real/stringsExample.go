package main

import (
	"log"
	"unicode/utf8"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func ReadMemory(ptr unsafe.Pointer, size uintptr) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))))
	}
	return out
}

type goString struct {
	elements []byte // underlying string bytes
	len      int    // number of bytes
}

// 文
func main() {
	s := []byte("Hello world")
	var stringExample = "Hello world"
	var anotherStringExample = "Hello world"
	var goString = goString{s, 11}

	sz := unsafe.Sizeof(stringExample)
	log.Println(sz) // 16

	log.Println(unsafe.Pointer(&stringExample))
	log.Println(unsafe.Pointer(&anotherStringExample))

	stringExample = anotherStringExample

	log.Println(unsafe.Pointer(&stringExample))
	log.Println(unsafe.Pointer(&anotherStringExample))

	n := unsafe.Pointer(&goString.elements[0])
	log.Println(ReadMemory(n, 11))
	log.Println(string(ReadMemory(n, 11)))

	log.Println(utf8.RuneLen('文')) //3
	buf := []byte{0, 0, 0}
	utf8.EncodeRune(buf, '文')

	r, _ := utf8.DecodeRune(buf)
	log.Printf("%q", r)
}
