package random

import (
	"errors"
	"math/rand"
)

var ErrRandRange = errors.New("the range of RandIntIn should be larger than 0")

// Generate random integer
func RandInt() int {
	return rand.Int()
}

// Generate random integer in range [l, r), return ErrRandRange if l is large or equal than r
func RandIntIn(l, r int) (int, error) {
	if l >= r {
		return 0, ErrRandRange
	}
	return rand.Intn(r-l) + l, nil
}
