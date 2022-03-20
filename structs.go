package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var contaDeWalter ContaCorrente = ContaCorrente{
		titular:       "Walter",
		numeroAgencia: 001,
		numeroConta:   123123,
		saldo:         2.87,
	}

	contaDeJunior := ContaCorrente{
		titular:       "Junior",
		numeroAgencia: 001,
		numeroConta:   123124,
		saldo:         12.59,
	}

	fmt.Println(contaDeWalter)
	fmt.Println(contaDeJunior)
}
