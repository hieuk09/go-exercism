package triangle

import "math"

type Kind struct {
	name string
}

var Equ = Kind{name: "equilateral"}
var Iso = Kind{name: "isosceles"}
var Sca = Kind{name: "scalene"}
var NaT = Kind{name: "not valid"}

func KindFromSides(a, b, c float64) Kind {
	switch {
	case IsInvalidNumber(a, b, c) || IsNegativeLength(a, b, c) || IsFailedTriangleInequality(a, b, c):
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}

func IsFailedTriangleInequality(a, b, c float64) bool {
	return (a+b <= c) || (a+c <= b) || (b+c <= a)
}

func IsNegativeLength(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0
}

func IsInvalidNumber(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
}
