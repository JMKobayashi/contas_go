package main

import (
	"fmt"

	"github.com/contas_go/clientes"
	"github.com/contas_go/contas"
)

func main() {
	contaDoJames := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "James",
		CPF:       "390.390.390-90",
		Profissao: "Desenvolvedor",
	}, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoJames.Depositar(1000)

	clienteMarjorie := clientes.Titular{Nome: "Marjorie", CPF: "123.456.789-10", Profissao: "Projetista"}

	contaDaMarjorie := contas.ContaCorrente{Titular: clienteMarjorie, NumeroAgencia: 123, NumeroConta: 12345}
	contaDaMarjorie.Depositar(1500)

	fmt.Println(contaDoJames, contaDaMarjorie)
}
