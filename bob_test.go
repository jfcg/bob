package bob

import (
	"sort"
	"testing"
)

// ls size, # of goroutines
const N, G uint = 18e8, 8

var (
	te *testing.T
	ls = make([]uint64, N)
	ch = make(chan bool, G/2)
	uc chan uint64
)

// ls[x:y] different neighbors?
func chk(x, y uint) {
	for i := x; i > y; i-- {
		if ls[i] == ls[i-1] {
			te.Fatal("bad Bob!")
		}
	}
	ch <- false
}

// fill ls with b's output, sort ls, cry on repeating elements
func tsb(b *Bob) {
	for i := N; i > 0; i-- {
		ls[i-1] = b.Squeez()
	}
	sort.Slice(ls, func(i, j int) bool { return ls[i] < ls[j] })

	x, st := N-1, (N+G/2)/G // start, length of one interval
	for i := G - 1; i > 0; i-- {
		go chk(x, x-st)
		x -= st
	}
	chk(x, 0)

	// wait for friends
	for i := G - 1; i > 0; i-- {
		<-ch
	}
}

// search x in ls
func sch() {
	for {
		x, ok := <-uc
		if !ok {
			break
		}

		p, q := uint(0), N
		for p < q {
			m := (p + q) / 2
			y := ls[m]
			if x == y {
				te.Fatal("bad Bob!")
			}

			if x < y {
				q = m
			} else {
				p = m + 1
			}
		}
	}
	ch <- false
}

// save original Bob, seek a's output in sorted list ls, then tsb(original)
func cmb(a *Bob) {
	b := *a
	uc = make(chan uint64, G-1)

	for i := G - 1; i > 0; i-- {
		go sch()
	}

	for i := N; i > 0; i-- {
		uc <- a.Squeez()
	}
	close(uc)

	// wait for friends
	for i := G - 1; i > 0; i-- {
		<-ch
	}

	tsb(&b)
}

func Test1(x *testing.T) {
	te = x
	var a, b Bob
	tsb(&b)
	a.Absorb(1)
	cmb(&a)
}
