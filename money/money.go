package money

type (
	Dollar struct {
		amount int
	}
)

func newDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return newDollar(d.amount * multiplier)
}

func (d *Dollar) Equals(p *Dollar) bool {
	return d.amount == p.amount
}
