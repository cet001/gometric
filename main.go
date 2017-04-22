// Package gometric is a parent package for various families of distance functions.
//
// String distance functions are in package gometric/strdist.
//
// Geographical distance functions are in package gometric/geodist.
package gometric

// Explicitly import the subpackages so that `go install` will "reach" them.
import (
	_ "github.com/cet001/gometric/strdist"
)
