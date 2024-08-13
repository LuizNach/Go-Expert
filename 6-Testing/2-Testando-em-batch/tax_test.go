package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	// ARRANGE
	amount := 500.00
	expected := 5.0

	// ACT
	result := CalculateTax(amount)

	// ASSERT
	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

}

func TestCalculateTaxBatch(t *testing.T) {

	// ARRANGE
	type CalcTax struct {
		amount, expected float64
	}

	table := []CalcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{10.0, 5.0},
		{2000.0, 10.0},
	}

	for _, v := range table {
		// ACT
		result := CalculateTax(v.amount)

		// ASSERT
		if result != v.expected {
			t.Errorf("Expected %f but got %f", v.expected, result)
		}
	}
}
