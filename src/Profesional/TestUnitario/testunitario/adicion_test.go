package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("El resutado es %d y se esperaba %d\n", total, 10)
	}
}
