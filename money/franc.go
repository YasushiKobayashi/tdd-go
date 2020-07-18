package money

type (
	Franc struct {
		amount int
	}
)

func newFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (d *Franc) Times(multiplier int) *Franc {
	return newFranc(d.amount * multiplier)
}

func (d *Franc) Equals(p *Franc) bool {
	return d.amount == p.amount
}
