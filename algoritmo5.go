package main

import "math"

func AverageWithPi(ptr *float64) {
	*ptr = (math.Pi + *ptr) / 2
}
