package strdist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevenshtein_Dist(t *testing.T) {
	dist := NewLevenshtein().Dist

	assert.Equal(t, 0, dist("", ""))
	assert.Equal(t, 0, dist("foo", "foo"))

	assert.Equal(t, 3, dist("foo", ""))
	assert.Equal(t, 3, dist("", "foo"))
	assert.Equal(t, 3, dist("foo", "bar"))

	assert.Equal(t, 3, dist("kitten", "sitting"))
	assert.Equal(t, 3, dist("saturday", "sunday"))
}

func Benchmark_Levenshtein_Dist(b *testing.B) {
	s1values := []string{"martha", "dixon", "apple", "constitution", "mississippi"}
	s2values := []string{"marhta", "dicksonx", "microsoft", "intervention", "misanthrope"}
	numValues := len(s1values)
	i := 0

	lev := NewLevenshtein()

	calcDist := func() {
		lev.Dist(s1values[i], s2values[i])
		i++
		if i == numValues {
			i = 0
		}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		calcDist()
	}
}
