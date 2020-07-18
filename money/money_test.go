package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := newDollar(5)
	five.times(2)

	if 10 != five.amount {
		t.Fatalf("amount is 10, but get %d", five.amount)
	}
}
