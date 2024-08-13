package tax

import "testing"

func Test_calculate_tax(t *testing.T) {
	amount := 500.00
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

}
