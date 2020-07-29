package random

import (
	"errors"
	"math/rand"
)

var ErrRandRange = errors.New("the range of RandIntIn should be larger than 0")

func RandInt() int {
	return rand.Int()
}

func RandIntIn(l, r int) (int, error) {
	if l == r {
		return 0, ErrRandRange
	}
	return rand.Intn(r-l) + l, nil
}
