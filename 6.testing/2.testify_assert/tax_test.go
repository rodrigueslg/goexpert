package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculeTax(0.0)
	assert.Error(t, err, "amount should be greater than 0")
	assert.Equal(t, 0.0, tax)

	tax, err = CalculeTax(1000.0)
	assert.Error(t, err, "amount should be greater than 1000 and less than 20000")
	assert.Equal(t, 10.0, tax)

	tax, err = CalculeTax(5.0)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, tax)
}
