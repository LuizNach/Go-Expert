package tax

func CalculateTax(amount float64) float64 {
	if amount == 0 {
		return 0
	} else if amount >= 1000.00 {
		return 10.0
	} else {
		return 5.0
	}
}
