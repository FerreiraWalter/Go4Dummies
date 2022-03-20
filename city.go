package main

import "fmt"

func main() {
	city, population, capital := showCityandPopulation()
	if capital {
		fmt.Println("Capital:", city, "Population:", population)
	} else {
		fmt.Println("City:", city, "Population:", population)
	}
}

func showCityandPopulation() (string, int, bool) {
	return "Ajacaru", 2003, true
}
