package main

import (
	"fmt"
	"poo/clientes"
	"poo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDeWalter := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Walter",
		CPF:       "123.123.123-12",
		Profissao: "Melhor dev do bairro dele",
	},
		NumeroAgencia: 0111,
		NumeroConta:   1234}

	clienteDoEstranho := clientes.Titular{Nome: "Estranho", CPF: "010101010", Profissao: "Segundo melhor dev"}
	contaDoEstranho := contas.ContaCorrente{NumeroAgencia: 0111, NumeroConta: 1235}

	fmt.Println(contaDeWalter)
	fmt.Println(clienteDoEstranho, contaDoEstranho)

	contaDoDouglas := contas.ContaPoupanca{}
	contaDoDouglas.Depositar(100)
	PagarBoleto(&contaDoDouglas, 60)

	fmt.Println(contaDoDouglas)
	fmt.Println(contaDoDouglas.ObterSaldo())
}
