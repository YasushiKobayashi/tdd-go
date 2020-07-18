package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := newDollar(5)

	product := five.Times(2)
	if 10 != product.amount {
		t.Fatalf("amount is 10, but get %d", five.amount)
	}

	product = five.Times(3)
	if 15 != product.amount {
		t.Fatalf("amount is 15, but get %d", five.amount)
	}
}

func TestEquality(t *testing.T) {
	five := newDollar(5)
	newFive := newDollar(5)
	if !five.Equals(newFive) {
		t.Fatalf("five %v equals %v", five, newFive)
	}

	newSix := newDollar(6)
	if five.Equals(newSix) {
		t.Fatalf("five %v not equals %v", five, newFive)
	}
}
