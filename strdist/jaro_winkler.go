package strdist

import (
	"github.com/cet001/mathext/ints"
)

const maxStringLen = 1024 * 4 // size of working space

// Calculates the Jaro string distance (similarity) score.
// Based on the Wikipedia description: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance.
type Jaro struct {
	s1matches []bool
	s2matches []bool
}

func NewJaro() *Jaro {
	return &Jaro{
		s1matches: make([]bool, maxStringLen),
		s2matches: make([]bool, maxStringLen),
	}
}

// Returns the Jaro distance score for s1 and s2.
//
// WARNING: This method is NOT threadsafe!
func (me *Jaro) Dist(s1, s2 string) float64 {
	lenS1, lenS2 := len(s1), len(s2)

	if lenS1 > maxStringLen {
		lenS1 = maxStringLen
		s1 = s1[:maxStringLen]
	}

	if lenS2 > maxStringLen {
		lenS2 = maxStringLen
		s2 = s2[:maxStringLen]
	}

	maxMatchDist := (ints.Max(lenS1, lenS2) / 2) - 1
	s1matches, s2matches := me.s1matches, me.s2matches

	// Clear the working space
	for i := 0; i < ints.Max(lenS1, lenS2); i++ {
		s1matches[i] = false
		s2matches[i] = false
	}

	// Count the matches and track which characters from each string matched.
	m := float64(0)
	for i := 0; i < lenS1; i++ {
		ch1 := s1[i]
		left, right := ints.Max(0, i-maxMatchDist), ints.Min(lenS2, i+maxMatchDist+1)
		for j := left; j < right; j++ {
			ch2 := s2[j]
			if ch1 == ch2 && !s2matches[j] {
				m++
				s1matches[i] = true
				s2matches[j] = true
				break
			}
		}
	}

	if m == 0 { // no matches
		return 0.0
	}

	// Calculate the number of full transpositions
	halfTranspositions := 0
	j := 0
	for i := 0; i < lenS1; i++ {
		if s1matches[i] {
			for !s2matches[j] {
				j++
			}
			if s1[i] != s2[j] {
				halfTranspositions++
			}
			j++
		}
	}
	t := float64(halfTranspositions / 2)

	return ((m / float64(lenS1)) + (m / float64(lenS2)) + ((m - t) / m)) / 3.0
}

// Calculates the Jaro-Winkler string distance (similarity) score.
type JaroWinkler struct {
	jaro        *Jaro
	prefixScale float64
}

func NewJaroWinkler() *JaroWinkler {
	return &JaroWinkler{
		jaro:        NewJaro(),
		prefixScale: 0.10,
	}
}

// Returns the Jaro-Winkler distance score for s1 and s2.
//
// WARNING: This method is NOT threadsafe!
func (me *JaroWinkler) Dist(s1, s2 string) float64 {
	lenS1, lenS2 := len(s1), len(s2)
	jaroDist := me.jaro.Dist(s1, s2)

	// Let p be the length of the largest common prefix of s1 and s2 (limit to 4 chars)
	p := 0
	for i := 0; i < lenS1; i++ {
		if i < lenS2 {
			if s1[i] != s2[i] {
				break
			}
			p++
			if p == 4 {
				break
			}
		}
	}

	return jaroDist + float64(p)*me.prefixScale*(1.0-jaroDist)
}
