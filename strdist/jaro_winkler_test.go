package strdist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func ExampleJaroWinkler() {
	jaroWinkler := NewJaroWinkler()
	fmt.Println(jaroWinkler.Dist("andrew", "andrew"))
	fmt.Println(jaroWinkler.Dist("martha", "marhta"))
	fmt.Println(jaroWinkler.Dist("jones", "johnson"))
	fmt.Println(jaroWinkler.Dist("foo", "bar"))
	// Output:
	// 1
	// 0.9611111111111111
	// 0.8323809523809523
	// 0
}

func TestJaroWinkler_Dist(t *testing.T) {
	jaroWinkler := NewJaroWinkler()

	// This function asserts that the expected distance score is returned by both
	// Dist(s1, s2) and Dist(s2, s1).
	assertDist := func(expected float64, s1, s2 string) {
		assert.Equal(t, expected, jaroWinkler.Dist(s1, s2), fmt.Sprintf("Dist('%v', '%v')", s1, s2))
		assert.Equal(t, expected, jaroWinkler.Dist(s2, s1), fmt.Sprintf("Dist('%v', '%v')", s2, s1))
	}

	assertDist(0.0, "", "")
	assertDist(0.0, "abc", "")
	assertDist(0.0, "", "abc")
	assertDist(0.0, "abc", "xyz")

	assertDist(1.0, "abc", "abc")
	assertDist(1.0, "takahashi", "takahashi")

	// See examples 1 and 2 from https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
	assertDist(0.9611111111111111, "martha", "marhta")
	assertDist(0.8133333333333332, "dixon", "dicksonx")

	// See examples from http://alias-i.com/lingpipe/docs/api/com/aliasi/spell/JaroWinklerDistance.html
	assertDist(0.8323809523809523, "jones", "johnson")
}

func TestJaroWinkler_Dist_limitCommonPrefixLength(t *testing.T) {
	jaroWinkler := NewJaroWinkler()
	d := jaroWinkler.Dist("aaaaaaaaaaaaaaaaaabbbbbbbbbbbbbb", "aaaaaaaaaaaaaaaaaacccccccccccccc")
	assert.True(t, d < 1.0)
}

func TestJaroWinkler_Dist_veryLongStrings(t *testing.T) {
	jaroWinkler := NewJaroWinkler()

	s1 := strings.Repeat("a", maxStringLen+1)
	s2 := "foo"

	assert.NotPanics(t, func() {
		jaroWinkler.Dist(s1, s2)
	})

	assert.NotPanics(t, func() {
		jaroWinkler.Dist(s2, s1)
	})
}

func Benchmark_JaroWinkler_Dist(b *testing.B) {
	s1values := []string{"martha", "dixon", "apple", "constitution", "mississippi"}
	s2values := []string{"marhta", "dicksonx", "microsoft", "intervention", "misanthrope"}
	numValues := len(s1values)
	i := 0

	jaroWinkler := NewJaroWinkler()

	calcDist := func() {
		jaroWinkler.Dist(s1values[i], s2values[i])
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
