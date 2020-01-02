package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	testes := []struct {
		num1 int
		num2 int
		resultado int
	} {
		{1,1,2},
		{1,2,3},
		{2,2,4},
		{2,3,5},
		{5,5,10},
	}

	for _, teste := range testes {
		t.Logf("Testando: %v + %v \n", teste.num1, teste.num2)
		caso := Soma(teste.num1, teste.num2)

		if caso != teste.resultado {
			t.Errorf("A soma esperada entre %v e %v é: %v e não %v", teste.num1, teste.num2, teste.resultado, caso)
		}
	}
}