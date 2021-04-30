// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

package triangle

import "math"

type Kind int

const (
	NaT = 0
	Equ = 1
	Iso = 2
	Sca = 3
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	if !(a > 0 && b > 0 && c > 0 &&
		a != pinf && b != pinf && c != pinf && a != ninf && b != ninf && c != ninf && a+b >= c &&
		b+c >= a && a+c >= b) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k

}
