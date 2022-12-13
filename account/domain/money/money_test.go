package money_test

import (
	"testing"

	"github.com/longyue0521/buckpal-in-go/account/domain/money"
	"github.com/stretchr/testify/assert"
)

func TestMoney(t *testing.T) {

	negOne := money.NewMoney(-1)
	zero := money.NewMoney(0)
	one := money.NewMoney(1)

	assert.False(t, negOne.IsPositiveOrZero())
	assert.True(t, zero.IsPositiveOrZero())
	assert.True(t, one.IsPositiveOrZero())

	assert.True(t, negOne.IsNegative())
	assert.False(t, zero.IsNegative())
	assert.False(t, one.IsNegative())

	assert.False(t, negOne.IsPositive())
	assert.False(t, zero.IsPositive())
	assert.True(t, one.IsPositive())

	assert.True(t, one.IsGreaterThanOrEqualTo(zero))
	assert.True(t, one.IsGreaterThanOrEqualTo(one))
	assert.False(t, zero.IsGreaterThanOrEqualTo(one))

	assert.True(t, one.IsGreaterThan(zero))
	assert.False(t, zero.IsGreaterThan(zero))
	assert.False(t, negOne.IsGreaterThan(zero))

	assert.True(t, money.Add(one, negOne).IsGreaterThanOrEqualTo(zero))

	assert.True(t, money.Subtract(zero, negOne).IsGreaterThanOrEqualTo(one))

	assert.True(t, one.Negate().IsGreaterThanOrEqualTo(negOne))
	assert.True(t, negOne.Negate().IsGreaterThanOrEqualTo(one))

	assert.True(t, one.Minus(one).IsGreaterThanOrEqualTo(zero))
	assert.True(t, one.Minus(zero).IsGreaterThanOrEqualTo(one))
	assert.True(t, zero.Minus(negOne).IsGreaterThanOrEqualTo(one))

	assert.True(t, one.Plus(negOne).IsGreaterThanOrEqualTo(zero))
	assert.True(t, one.Plus(zero).IsGreaterThanOrEqualTo(one))
	assert.True(t, zero.Plus(negOne).IsGreaterThanOrEqualTo(negOne))
}
