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
	
	var firstName string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("How many tickets would you like to have: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
}
