package main

import (
	"log"
	"reflect"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type User struct {
	name string
	age  int
}

func main() {
	alex := User{}
	log.Println(alex.age) //0
	alexP := &alex

	log.Println(alexP)

	var Worker = struct {
		User
		salary int
	}{
		User:   alex,
		salary: 100000,
	}

	var AnotherWorker = struct {
		User
		salary int
	}{
		User: struct {
			name string
			age  int
		}{
			"",
			0,
		},
		salary: 100000,
	}

	log.Println(Worker.salary)
	log.Println(Worker.name)
	log.Println(Worker.age)

	log.Println(AnotherWorker == Worker) // true

	a := struct {
		name string
		age  int
	}{"", 0}
	b := User{
		name: "",
		age:  0,
	}
	log.Println(reflect.DeepEqual(a, alex))               // false
	log.Println(reflect.DeepEqual(b, alex))               // true
	log.Println(reflect.DeepEqual(AnotherWorker, Worker)) // true
}
