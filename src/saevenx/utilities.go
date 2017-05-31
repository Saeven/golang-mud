package saevenx

import (
	"time"
	"math/rand"
)

/**
 * Stuff that should probably exist in Go (in my opinion)
 */

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func random_int(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}