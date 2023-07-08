package symbol_test

import (
	"testing"

	"github.com/fogfish/it/v2"
	"github.com/kshard/xsd/symbol"
)

func TestSymbols(t *testing.T) {
	symbols := symbol.New()

	a, err := symbols.SymbolOf("a")
	it.Then(t).Should(
		it.Nil(err),
		it.Greater(a, 0),
	)

	b, err := symbols.SymbolOf("a")
	it.Then(t).Should(
		it.Nil(err),
		it.Equal(a, b),
		it.Equal(symbols.ValueOf(a), "a"),
		it.Equal(symbols.ValueOf(b), "a"),
	)
}

func TestNotFound(t *testing.T) {
	symbols := symbol.New()

	v := symbols.ValueOf(123)
	it.Then(t).Should(
		it.Equal(v, ":123"),
	)
}

func TestCollision(t *testing.T) {
	symbols := symbol.New()

	for _, in := range [][]string{
		{"8yn0iYCKYHlIj4-BwPqk", "GReLUrM4wMqfg9yzV3KQ"},
		{"gMPflVXtwGDXbIhP73TX", "LtHf1prlU1bCeYZEdqWf"},
		{"pFuM83THhM-Qw8FI5FKo", ".jPx7rOtTDteKAwvfOEo"},
		{"7mohtcOFVz", "c1E51sSEyx"},
		{"6a5x-VbtXk", "f_2k7GG-4v"},
	} {
		a, err := symbols.SymbolOf(in[0])
		it.Then(t).Should(
			it.Nil(err),
			it.Greater(a, 0),
		)

		b, err := symbols.SymbolOf(in[1])
		it.Then(t).Should(
			it.Nil(err),
			it.Greater(a, 0),
		)

		it.Then(t).ShouldNot(
			it.Equal(a, b),
		).Should(
			it.Equal(symbols.ValueOf(a), in[0]),
			it.Equal(symbols.ValueOf(b), in[1]),
		)
	}
}

// ---------------------------------------------------------------

// go test -fuzz=FuzzSymbolOf
func FuzzSymbolOf(f *testing.F) {
	set := symbol.New()
	f.Add("abc")

	f.Fuzz(func(t *testing.T, el string) {
		_, err := set.SymbolOf(el)
		if err != nil {
			t.Errorf("%s", err)
		}
	})
}
