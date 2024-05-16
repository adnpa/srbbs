package lib

import (
	"math"
)

func Sign(num float64) int {
	if num > 0 {
		return 1
	} else if num == 0 {
		return 0
	} else {
		return -1
	}
}

func Threshold(x float64) float64 {
	if math.Abs(x) >= 1 {
		return math.Abs(x)
	} else {
		return 1
	}
}
