package main

import (
	"fmt"

	"github.com/danfmsouza/banco/contas"
)

func main() {
	contaDoLoxa := contas.ContaPoupanca{}
	contaDoLoxa.Depositar(100)
	contaDoLoxa.Sacar(55)

	fmt.Println(contaDoLoxa.ObterSaldo())

	PagarBoleto(&contaDoLoxa, 10)
	fmt.Println(contaDoLoxa.ObterSaldo())

	contaDoDaniel := contas.ContaCorrente{}
	contaDoDaniel.Depositar(500)

	PagarBoleto(&contaDoDaniel, 499)
	fmt.Println(contaDoDaniel.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
