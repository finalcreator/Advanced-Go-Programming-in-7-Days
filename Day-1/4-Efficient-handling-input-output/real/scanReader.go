package main

import (
	"bufio"
	"fmt"
	"io"
)

type myReader struct {
	text string
}

func newMyReader(text string) *myReader {
	return &myReader{text: text}
}

func (r *myReader) Read(p []byte) (int, error) {
	fmt.Println("\n------------")
	buf := make([]byte, len(r.text))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.text[i] | 0x20
	}

	n := copy(p, buf)
	return n, io.EOF
}

func main() {
	// Example of implementing a custom io.Reader
	r := newMyReader("ALL CAPITALS")
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

}
