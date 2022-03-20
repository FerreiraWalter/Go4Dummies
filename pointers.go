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

	var contaDeWalter2 ContaCorrente = ContaCorrente{
		titular:       "Walter",
		numeroAgencia: 001,
		numeroConta:   123123,
		saldo:         2.87,
	}

	fmt.Println("Contas walters:", contaDeWalter == contaDeWalter2)

	contaDeJunior := ContaCorrente{
		titular:       "Junior",
		numeroAgencia: 001,
		numeroConta:   123124,
		saldo:         12.59,
	}

	contaDeJunior2 := ContaCorrente{
		titular:       "Junior",
		numeroAgencia: 001,
		numeroConta:   123124,
		saldo:         12.59,
	}

	fmt.Println("Contas juniors:", contaDeJunior == contaDeJunior2)

	var contaDaMaria *ContaCorrente
	contaDaMaria = new(ContaCorrente)
	contaDaMaria.titular = "Maria"
	contaDaMaria.numeroAgencia = 001
	contaDaMaria.numeroConta = 123125
	contaDaMaria.saldo = 504.10

	var contaDaMaria2 *ContaCorrente
	contaDaMaria2 = new(ContaCorrente)
	contaDaMaria2.titular = "Maria"
	contaDaMaria2.numeroAgencia = 001
	contaDaMaria2.numeroConta = 123125
	contaDaMaria2.saldo = 504.10

	fmt.Println("Contas marias:", contaDaMaria == contaDaMaria2)
	fmt.Println("Contas marias *:", *contaDaMaria == *contaDaMaria2)
}
