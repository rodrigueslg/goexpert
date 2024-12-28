package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculeTax(amount)

	if result != expected {
		t.Errorf("Expected: %f, got: %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type spec struct {
		amount   float64
		expected float64
	}

	table := []spec{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, row := range table {
		res := CalculeTax(row.amount)
		if res != row.expected {
			t.Errorf("Expected: %f, got: %f", row.expected, res)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculeTax(500.0)
	}
}

func BenchmarkCalculateTaxDelayed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculeTaxDelayed(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, -1000.0, -1500.0, -1501.0}
	for _, s := range seed {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		res := CalculeTax(amount)
		if amount <= 0 && res != 0 {
			t.Errorf("Expected: 0, got: %f", res)
		}
		if amount > 20000 && res != 20 {
			t.Errorf("Expected: 20, got: %f", res)
		}
	})
}
