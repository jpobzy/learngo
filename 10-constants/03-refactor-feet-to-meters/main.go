// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

// EXERCISE
//
// This program uses magic numbers
// Refactor it and use named constants instead

func main() {
	
	const (
		feetInMeters = 0.3048
		feetInYards = feetInMeters / 0.9144
	)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)


	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
