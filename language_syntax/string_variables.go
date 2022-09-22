package main

import (
	"fmt"
)

var (
	Name1 string = "global1"
	Name2 string = "global2"
)

var String1, String2 string = "local1", "local2"

func main() {
	var name1 string
	var name2 string = "CloudAcademy"
	name3 := "DevOps 2022"
	name4 := name2 + " " + name3
	fmt.Println("This is a test statement.")
	fmt.Println(Name1)
	fmt.Println(Name2)
	fmt.Println(String1)
	fmt.Println(String2)
	String1 = "local_changed1"
	fmt.Println(String1)
	var Name1 string = "global_changed1"
	fmt.Println(Name1)
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Printf("%v ---- %v\n", Name1, Name2)
	fmt.Println(name2, name3, Name1, Name2)
}
