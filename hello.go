package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Walter"
	second_name := "Junior"
	var age int = 20
	var version float32 = 1.18

	fmt.Println("Hello", name, second_name, "your age is", age)
	fmt.Println("Version:", version)
	fmt.Println("Var type:", reflect.TypeOf(name))
}
