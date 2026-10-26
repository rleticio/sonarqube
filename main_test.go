package main

import "testing"

func TestCalcularTotal(t *testing.T) {
    total := calcularTotal(50, 3)
    esperado := 150.0

    if total != esperado {
        t.Errorf("calcularTotal(50, 3) = %v; esperado %v", total, esperado)
    }
}

func TestExecutar(t *testing.T) {
    total := executar()
    esperado := 150.0

    if total != esperado {
        t.Errorf("executar() = %v; esperado %v", total, esperado)
    }
}

func TestMainFunction(t *testing.T) {
    main()
}

func TestCalcularTotalPrecoNegativo(t *testing.T) {
    total := calcularTotal(-50, 3)
    esperado := 0.0

    if total != esperado {
        t.Errorf(
            "calcularTotal(-50, 3) = %v; esperado %v",
            total,
            esperado,
        )
    }
}

func TestVerificarPreco(t *testing.T) {
    verificarPreco(100)
}

func TestCalcularDobroETriplo(t *testing.T) {
    if calcularDobro(5) != 10 {
        t.Error("calcularDobro(5) deveria retornar 10")
    }

    if calcularTriplo(5) != 15 {
        t.Error("calcularTriplo(5) deveria retornar 15")
    }
}

func TestCalcularDesconto(t *testing.T) {
    resultado := calcularDesconto(100)
    if resultado != 90 {
        t.Errorf("esperado 90, recebido %.2f", resultado)
    }
}

func TestVerificarAprovacao(t *testing.T) {
    if verificarAprovacao(8) != "Aprovado" {
        t.Error("nota 8 deveria ser aprovada")
    }

    if verificarAprovacao(5) != "Reprovado" {
        t.Error("nota 5 deveria ser reprovada")
    }
}

func TestTestarRegraSonar(t *testing.T) {
    if !testarRegraSonar(true) {
        t.Error("esperado true para entrada true")
    }

    if testarRegraSonar(false) {
        t.Error("esperado false para entrada false")
    }
}











