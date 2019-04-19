package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SliceOfStringConvSliceOfAscii(data string) {
	numbers := strings.Split(data, ",")
	//_ := make([]int, len(numbers))
	for index, char := range numbers {
		fmt.Printf("char(%), index(%d)\n", char, index)
	}

}

func main() {
	numberInput := "1,2,3,10,11,12,100,200,8888,9999"
	SliceOfStringConvSliceOfAscii(numberInput)

	os.Exit(0)

	intScanner := bufio.NewScanner(strings.NewReader(numberInput))

	// custom split by comma function
	splitByComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Print("\n----- splitByComma -----\n")
		fmt.Printf("data(%v)\tatEOF(%t)\n", data, atEOF)

		if atEOF && len(data) == 0 {
			fmt.Println("atEOF && len(data) == 0")
			fmt.Println("return 0, nil, nil")
			return 0, nil, nil
		}

		if i := strings.IndexRune(string(data), ','); i >= 0 {
			fmt.Printf("i:= strings.IndexRune(string(data),',') i=%v \n", i)
			fmt.Printf("return i + 1 :%v , data[0:i] : %v, nil\n", i+1, data[0:i])
			return i + 1, data[0:i], nil
		}

		if atEOF {
			fmt.Println("atEOF=true")
			fmt.Printf("return len(data):%v, data:%v, nil\n", len(data), data)
			return len(data), data, nil
		}
		return
	}

	intScanner.Split(splitByComma)

	// Scan 2 bytes at the time
	buf := make([]byte, 2)
	intScanner.Buffer(buf, bufio.MaxScanTokenSize)
	for intScanner.Scan() {
		fmt.Printf("scan value: %s\n", intScanner.Text())
	}
}
