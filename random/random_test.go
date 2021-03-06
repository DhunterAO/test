package random

import "testing"

func TestRandInt(t *testing.T) {
	if randInt := RandInt(); randInt < 0 {
		t.Error("the random int should be non-negative")
	}
}

func TestRandIntIn(t *testing.T) {
	l := 1
	r := 100
	if rand, err := RandIntIn(l, r); err != nil || rand >= r || rand < l {
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Errorf("the random int should be in [%d, %d)", l, r)
		}
	}
}

func TestRandIntInErr(t *testing.T) {
	l := RandInt()
	r := l
	if rand, err := RandIntIn(l, r); rand != 0 || err != ErrRandRange {
		t.Error("RandIntIn should return 0 and ErrRandRange")
	}
}
