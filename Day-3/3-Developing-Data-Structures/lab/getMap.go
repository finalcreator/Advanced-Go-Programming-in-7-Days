package main

import "fmt"

func main() {
	table := make(map[string]string)
	table["a"] = "apple"
	table["b"] = "banana"
	table["c"] = "coco"
	a, s := table["b"]

	fmt.Println(a, s)
}
