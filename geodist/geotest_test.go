package geodist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApproxDist(t *testing.T) {
	// Various places in/around the Mission and Castro
	taraIndianCuisine := Coord{37.765492, -122.431229}
	blackbirdBar := Coord{37.767487, -122.429633}
	madDogInTheFog := Coord{37.772528, -122.431030}
	spinCityCoffee := Coord{37.749287, -122.429181}
	civicCenterBart := Coord{37.780396, -122.414135}

	const delta = 0.2 // maximum allowable error is 20 meters (0.2Km)

	assert.InDelta(t, 0.27, ApproxDist(taraIndianCuisine, blackbirdBar), delta)
	assert.InDelta(t, 0.85, ApproxDist(taraIndianCuisine, madDogInTheFog), delta)
	assert.InDelta(t, 2.4, ApproxDist(taraIndianCuisine, civicCenterBart), delta)
	assert.InDelta(t, 1.81, ApproxDist(taraIndianCuisine, spinCityCoffee), delta)
}
