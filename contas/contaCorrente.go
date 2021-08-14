package contas

import "github.com/danfmsouza/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado!"
	} else {
		return "saldo Insficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0 && (valorDoDeposito+c.saldo) > c.saldo
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado"
	} else {
		return "Depósito não realizado"
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valorDaTransferencia > 0 && valorDaTransferencia < c.saldo
	if podeTransferir {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
