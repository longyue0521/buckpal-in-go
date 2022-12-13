package money

import "math/big"

var zero = big.NewInt(0)

type Money struct {
	amount *big.Int
}

func NewMoney(value int64) *Money {
	return &Money{
		amount: big.NewInt(value),
	}
}

func Add(a *Money, b *Money) *Money {
	return NewMoney(a.amount.Int64() + b.amount.Int64())
}

func Subtract(a *Money, b *Money) *Money {
	return NewMoney(a.amount.Int64() - b.amount.Int64())
}

func (m *Money) IsPositiveOrZero() bool {
	return m.amount.Cmp(zero) >= 0
}

func (m *Money) IsNegative() bool {
	return m.amount.Cmp(zero) < 0
}

func (m *Money) IsPositive() bool {
	return m.amount.Cmp(zero) > 0
}

func (m *Money) IsGreaterThanOrEqualTo(n *Money) bool {
	return m.amount.Cmp(n.amount) >= 0
}

func (m *Money) IsGreaterThan(n *Money) bool {
	return m.amount.Cmp(n.amount) > 0
}

func (m *Money) Negate() *Money {
	return NewMoney(0 - m.amount.Int64())
}

func (m *Money) Minus(n *Money) *Money {
	return NewMoney(m.amount.Int64() - n.amount.Int64())
}

func (m *Money) Plus(n *Money) *Money {
	return NewMoney(m.amount.Int64() + n.amount.Int64())
}
