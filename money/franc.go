package money

type (
	Franc struct {
		*Money
	}
)

func newFranc(amount int) *Franc {
	return &Franc{
		Money: &Money{
			amount: amount,
		},
	}
}

func (d *Franc) Times(multiplier int) MoneyInterface {
	return newFranc(d.amount * multiplier)
}
