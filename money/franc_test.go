package money

import "testing"

func TestFrancMultiplication(t *testing.T) {
	five := newFranc(5)

	two := five.Times(2)
	newTwo := newFranc(10)
	if !two.Equals(newTwo.Money) {
		t.Fatalf("want %#v, but %#v", two, newTwo)
	}

	three := five.Times(3)
	newThree := newFranc(15)
	if !three.Equals(newThree.Money) {
		t.Fatalf("want %#v, but %#v", three, newThree)
	}
}
