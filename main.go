package main

import "fmt"

func calcularTotal(preco float64, quantidade int) float64 {
    if preco < 0 {
        return 0
    }

    return preco * float64(quantidade)
}

func executar() float64 {
    total := calcularTotal(50, 3)
    fmt.Println("Total:", total)
    return total
}

func main() {
    executar()
}

func precoEhValido(preco float64) bool {
    return preco >= 0
}

func verificarPreco(preco float64) {
    if precoEhValido(preco) == true {
        fmt.Println("Preço válido")
    }
}

func testarRegraSonar(valor bool) bool {
    //return valor && true
	return valor
}

func calcularDesconto(preco float64) float64 {
    desconto := 0.10
    return preco - (preco * desconto)
}

func verificarAprovacao(nota int) string {
    if nota >= 7 {
        return "Aprovado"
    } else {
        return "Reprovado"
    }
}

func calcularDobro(valor int) int {
    resultado := valor * 2
    return resultado
}

func calcularTriplo(valor int) int {
    resultado := valor * 3
    return resultado
}



