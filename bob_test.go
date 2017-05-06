package bob

import (
	"sort"
	"testing"
)

const N int = 4e8

// fill ls with b's output, sort ls, cry on repeating elements
func tsb(b *Bob, ls []uint64, t *testing.T) {
	for i := N - 1; i >= 0; i-- {
		ls[i] = b.Squeez()
	}
	sort.Slice(ls, func(i, j int) bool { return ls[i] < ls[j] })

	for i := N - 1; i > 0; i-- {
		if ls[i] == ls[i-1] {
			t.Fatal("bad Bob!")
		}
	}
}

// save original Bob, seek a's output in sorted list ls, then tsb(original)
func cmb(a *Bob, ls []uint64, t *testing.T) {
	b := *a
	for i := N - 1; i >= 0; i-- {
		x := a.Squeez()

		p, q := 0, N
		for p < q {
			m := (p + q) / 2
			y := ls[m]
			if x == y {
				t.Fatal("bad Bob!")
			}
			if x < y {
				q = m
			} else {
				p = m + 1
			}
		}
	}

	tsb(&b, ls, t)
}

func Test1(t *testing.T) {
	var a, b Bob
	ls := make([]uint64, N)
	tsb(&b, ls, t)
	a.Absorb(1)
	cmb(&a, ls, t)
}
