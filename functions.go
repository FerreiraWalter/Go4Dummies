package main

import (
	"fmt"
	"os"
)

func main() {
	showIntroduction()

	showMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying logs")
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Error")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Walter"
	version := 1.18

	fmt.Println("Hello", name)
	fmt.Println("Program version:", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("your choice was", command)

	return command
}
