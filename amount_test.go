package amount

import "testing"

func TestAmount(t *testing.T) {
	testData := []struct {
		amount    float64
		formatted string
	}{
		{
			amount:    0.5,
			formatted: "0,50",
		},
		{
			amount:    50,
			formatted: "50,00",
		},
		{
			amount:    123.4,
			formatted: "123,40",
		},
		{
			amount:    1234.56,
			formatted: "1.234,56",
		},
		{
			amount:    123456789.99,
			formatted: "123.456.789,99",
		},
	}

	for i, td := range testData {
		res := Amount(td.amount)
		if res != td.formatted {
			t.Fatalf("test %d failed, expected: %v got %v", i, td.formatted, res)
		}
	}
}
