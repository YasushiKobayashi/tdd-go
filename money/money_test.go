package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := newDollar(5)

	product := five.times(2)
	if 10 != product.amount {
		t.Fatalf("amount is 10, but get %d", five.amount)
	}

	product = five.times(3)
	if 15 != product.amount {
		t.Fatalf("amount is 15, but get %d", five.amount)
	}
}
