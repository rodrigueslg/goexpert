package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculeTax(0.0)
	assert.Equal(t, 0.0, tax)

	tax = CalculeTax(1000.0)
	assert.Equal(t, 10.0, tax)

	tax = CalculeTax(5.0)
	assert.Equal(t, 5.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repo := &TaxRepositoryMock{}
	repo.On("SaveTax", 10.0).Return(nil)
	repo.On("SaveTax", 0.0).Return(errors.New("amount should be greater than 0"))

	err := CalculateTaxAndSave(1000.0, repo)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repo)
	assert.Error(t, err, "amount should be greater than 0")

	repo.AssertExpectations(t)
}
