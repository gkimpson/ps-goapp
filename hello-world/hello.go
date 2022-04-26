package main

import (
	"fmt"
	"reflect"
)

var (
	name, course = "Gavin K", "Getting started with Go"
	module, clip = 4, 2
)

func main() {
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clup are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
}
