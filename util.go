package main

import (
	"math"
)

func NearlyEqual(a, b float64, epsilon ...float64) bool {
	defaultEpsilon := 1e-9

	// Use the provided epsilon or the default value if not provided
	var eps float64
	if len(epsilon) > 0 {
		eps = epsilon[0]
	} else {
		eps = defaultEpsilon
	}

	// already equal?
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	if a == 0.0 || b == 0.0 || diff < math.SmallestNonzeroFloat64 {
		return diff < eps*math.SmallestNonzeroFloat64
	}

	return diff/(math.Abs(a)+math.Abs(b)) < eps
}
