package tax

import (
	"testing"
)

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

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {

	seed := []float64{-1, -2, 0, 500.0, 1000.0, 1001.0, 2500.0}

	for _, value := range seed {
		f.Add(value)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount > 0 && amount < 1000.0 && result != 5 {
			t.Errorf("Received %f but expected 5", result)
		}

		if amount >= 1000.0 && amount < 20000.0 && result != 10 {
			t.Errorf("Received %f but expected 10", result)
		}

		if amount >= 20000.0 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	},
	)

}
