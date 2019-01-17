package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlusDecCoin(t *testing.T) {
	decCoinA1 := DecCoin{"A", NewDecWithPrec(11, 1)}
	decCoinA2 := DecCoin{"A", NewDecWithPrec(22, 1)}
	decCoinB1 := DecCoin{"B", NewDecWithPrec(11, 1)}

	// regular add
	res := decCoinA1.Plus(decCoinA1)
	require.Equal(t, decCoinA2, res, "sum of coins is incorrect")

	// bad denom add
	assert.Panics(t, func() {
		decCoinA1.Plus(decCoinB1)
	}, "expected panic on sum of different denoms")

}

func TestPlusDecCoins(t *testing.T) {
	one := NewDec(1)
	zero := NewDec(0)
	negone := NewDec(-1)
	two := NewDec(2)

	cases := []struct {
		inputOne DecCoins
		inputTwo DecCoins
		expected DecCoins
	}{
		{DecCoins{{"A", one}, {"B", one}}, DecCoins{{"A", one}, {"B", one}}, DecCoins{{"A", two}, {"B", two}}},
		{DecCoins{{"A", zero}, {"B", one}}, DecCoins{{"A", zero}, {"B", zero}}, DecCoins{{"B", one}}},
		{DecCoins{{"A", zero}, {"B", zero}}, DecCoins{{"A", zero}, {"B", zero}}, DecCoins(nil)},
		{DecCoins{{"A", one}, {"B", zero}}, DecCoins{{"A", negone}, {"B", zero}}, DecCoins(nil)},
		{DecCoins{{"A", negone}, {"B", zero}}, DecCoins{{"A", zero}, {"B", zero}}, DecCoins{{"A", negone}}},
	}

	for tcIndex, tc := range cases {
		res := tc.inputOne.Plus(tc.inputTwo)
		require.Equal(t, tc.expected, res, "sum of coins is incorrect, tc #%d", tcIndex)
	}
}