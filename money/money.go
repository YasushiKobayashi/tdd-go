package money

type (
	Money struct {
		amount int
	}

	MoneyInterface interface {
		Times(multiplier int) MoneyInterface
		Equals(p *Money) bool
	}
)

func (m *Money) Equals(p *Money) bool {
	return m.amount == p.amount
}
