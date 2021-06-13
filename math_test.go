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
