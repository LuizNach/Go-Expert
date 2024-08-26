package tax

import "time"

func CalculateTax(amount float64) float64 {

	if amount <= 0 {
		return 0
	}

	if amount >= 20000.0 {
		return 20.0
	}

	if amount >= 1000.0 {
		return 10.0
	}

	return 5.0
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	} else if amount >= 1000.00 {
		return 10.0
	} else {
		return 5.0
	}
}
