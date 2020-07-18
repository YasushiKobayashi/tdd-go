package money

type (
	Money struct {
		amount int
	}
)

func (m *Money) Equals(p *Money) bool {
	return m.amount == p.amount
}
