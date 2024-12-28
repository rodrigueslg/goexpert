package tax

import "errors"

func CalculeTax(amount float64) (float64, error) {
	if amount <= 0.0 {
		return 0.0, errors.New("amount should be greater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, errors.New("amount should be greater than 1000 and less than 20000")
	}
	if amount >= 20000 {
		return 20.0, errors.New("amount should be greater than 20000")
	}
	return 5.0, nil
}
