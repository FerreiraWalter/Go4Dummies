package main

import (
	"fmt"
)

func main() {
	name := "Walter"

	fmt.Println("Hello", name)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var command int
	fmt.Scanf("%d", &command)
	fmt.Println("your choice was", command)

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Displaying logs")
	// } else if command == 0 {
	// 	fmt.Println("Leaving the program")
	// } else {
	// 	fmt.Println("???")
	// }

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying logs")
	case 0:
		fmt.Println("Leaving the program")
	default:
		fmt.Println("???")
	}
}
