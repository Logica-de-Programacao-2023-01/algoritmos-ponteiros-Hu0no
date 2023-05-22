package main

import "fmt"

func PrimeVerification(x int) bool {
	if x == 0 || x == 1 {
		return false
	} else {
		z := 1
		qtd_divisores := 0
		for z <= x {
			if x%z == 0 {
				qtd_divisores += 1
			}
			z += 1
		}
		if qtd_divisores == 2 {
			return true
		} else {
			return false
		}
	}
}
func NumberOfPrimes(ptr *[]int, x int) {
	y := 2
	for len(*ptr) <= x {
		check := PrimeVerification(y)
		if check == true {
			*ptr = append(*ptr, y)
		}
		y++
	}
}
func main() {
	var x int = 6
	var slice []int
	var ptr *[]int = &slice
	NumberOfPrimes(ptr, x)
	fmt.Print(*ptr)
}
