package money

type (
	Dollar struct {
		*Money
	}
)

func newDollar(amount int) *Dollar {
	return &Dollar{
		Money: &Money{
			amount: amount,
		},
	}
}

func (d *Dollar) Times(multiplier int) MoneyInterface {
	return newDollar(d.amount * multiplier)
}
