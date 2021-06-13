package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido! Resultado: %d. Valor esperado: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := Sub(30, 15)

	if total != 15 {
		t.Errorf("Resultado da sub é inválido! Resultado: %d. Valor esperado: %d", total, 15)
	}
}

func TestTimes(t *testing.T) {
	total := Times(5, 5)

	if total != 25 {
		t.Errorf("Resultado da sub é inválido! Resultado: %d. Valor esperado: %d", total, 25)
	}
}

func TestDivides(t *testing.T) {
	total := Divide(10, 2)

	if total != 5 {
		t.Errorf("Resultado da sub é inválido! Resultado: %d. Valor esperado: %d", total, 5)
	}
}

func TestPower(t *testing.T) {
	total := Power(2, 2)

	if total != 4 {
		t.Errorf("Resultado da sub é inválido! Resultado: %d. Valor esperado: %d", total, 4)
	}
}
