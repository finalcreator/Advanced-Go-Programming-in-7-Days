package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type UserType reflect.Type

type User struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User: %v, %v", u.FirstName, u.LastName)
}

func main() {
	alex := User{}
	userType := reflect.TypeOf(alex)

	//log.Println(userType.Elem()) // panics
	log.Println(userType.NumField())             // 3
	log.Println(userType.Comparable())           // true
	log.Println(userType.Kind())                 // struct
	log.Println(userType.NumMethod())            // 1, Value vs Ref receiver matters
	log.Println(userType.MethodByName("String")) // case matters

	// Create slices via reflection
	intSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 0)
	log.Println(reflect.TypeOf(intSlice))
	log.Println(intSlice)
	intSlice = reflect.Append(intSlice, reflect.ValueOf(1))
	log.Println(intSlice) // [1]

	intArrayType := reflect.ArrayOf(5, reflect.TypeOf(0))
	intArray := reflect.New(intArrayType)
	log.Println(intArray) // &[0 0 0 0 0]

	var n = []int{1, 2, 3}
	var p = reflect.ValueOf(&n)
	log.Println(p)
	log.Println(reflect.TypeOf(p))
	log.Println(p.CanSet())  // false
	log.Println(p.CanAddr()) // false
	var nv = p.Elem()
	log.Println(reflect.TypeOf(nv))
	log.Println(nv.CanSet())  // true
	log.Println(nv.CanAddr()) // true
}
