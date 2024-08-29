package tax

import (
	"errors"
)

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {

	tax, err := CalculateTax(amount)

	if err != nil {
		println("Error!")
	}

	return repository.SaveTax(tax)

}

func CalculateTax(amount float64) (float64, error) {

	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0.0")
	}

	if amount >= 20000.0 {
		return 20.0, nil
	}

	if amount >= 1000.0 {
		return 10.0, nil
	}

	return 5.0, nil
}
