package money

import "testing"

func TestMoneyMultiplication(t *testing.T) {
	five := newDollar(5)

	two := five.Times(2)
	newTwo := newDollar(10)
	if !two.Equals(newTwo.Money) {
		t.Fatalf("%#v is not equals %#v", two, newTwo)
	}

	three := five.Times(3)
	newThree := newDollar(15)
	if !three.Equals(newThree.Money) {
		t.Fatalf("%#v is not equals %#v", three, newThree)
	}
}

func TestMoneyEquality(t *testing.T) {
	five := newDollar(5)
	newFive := newDollar(5)
	if !five.Equals(newFive.Money) {
		t.Fatalf("five %#v equals %#v", five, newFive)
	}

	newSix := newDollar(6)
	if five.Equals(newSix.Money) {
		t.Fatalf("five %#v not equals %#v", five, newFive)
	}
}
