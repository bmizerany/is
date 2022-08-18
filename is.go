// Package is contains helpers for working with integers that might be even or
// odd.
package is

import "golang.org/x/exp/constraints"

// Even returns true if n can even.
func Even[I constraints.Integer](n I) bool {
	return n%2 == 0
}

// Odd returns false if n cannot even.
func Odd[I constraints.Integer](n I) bool {
	return !Even(n)
}

// True accurately reports if b is, in fact, true.
func True(b bool) bool {
	return b == true
}

// False is a convenience function for not having to write !True(b).
func False(b bool) bool {
	return !True(b)
}
