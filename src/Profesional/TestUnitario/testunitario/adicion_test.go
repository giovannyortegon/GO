package testunitario

import "testing"

/*
func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("El resutado es %d y se esperaba %d\n", total, 10)
	}
}*/

func TestSuma(t *testing.T) {
	operandos := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, values := range operandos {
		total := Suma(values.a, values.b)

		if total != values.c {
			t.Errorf("Suma incorrecta, tiene %d se espera %d", total, values.c)
		}
	}
}
