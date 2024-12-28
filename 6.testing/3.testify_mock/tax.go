package tax

type Repository interface {
	SaveTax(float64) error
}

func CalculateTaxAndSave(amout float64, repo Repository) error {
	tax := CalculeTax(amout)
	return repo.SaveTax(tax)
}

func CalculeTax(amount float64) float64 {
	if amount <= 0.0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}
