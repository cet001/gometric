package geodist

import (
	"math"
)

// A geographical coordinate that represents a specific point on the earth.
type Coord struct {
	Lat float32
	Lng float32
}

// Returns the _approximate_ distance *in kilometers) between coordinates c1 and
// c2.  This distance function is much faster than, say, the Haversine fomrula
// or the spherical law of cosines, as it does not take into account the
// curvature of the earth.  Instead, this function trivially takes the Euclidian
// distance between the two points.  Because of this, the accuracy becomes worse
// the farther apart coordinates c1 and c2 become.
//
// This function is practical in situations where you need to calculat the
// rough distance between two "relatively close" points on a map (e.g. 2 houses
// located in the same neighborhood or city).
//
// See http://jonisalonen.com/2014/computing-distance-between-coordinates-can-be-simple-and-fast/
func ApproxDist(c1, c2 Coord) float64 {
	const degreeLengthKm float64 = 110.25 // length of 1 degree latitude at the equator
	x := float64(c1.Lat - c2.Lat)
	y := float64(c1.Lng-c2.Lng) * math.Cos(float64(c2.Lat))
	return degreeLengthKm * math.Sqrt(x*x+y*y)
}
