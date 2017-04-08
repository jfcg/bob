/*	Bob is a little sponge ;)

	Author: Serhat Sevki Dincer jfcgaussATgmail
*/

package bob

func chi1(x *[5]uint64) { // non-linear
	y := *x
	x[0] ^= y[1] &^ y[4]
	x[1] ^= y[2] &^ y[0]
	x[2] ^= y[3] &^ y[1]
	x[3] ^= y[4] &^ y[2]
	x[4] ^= y[0] &^ y[3]
}

func chi2(x *[5]uint64) {
	y := *x
	x[0] ^= y[3] &^ y[2]
	x[1] ^= y[4] &^ y[3]
	x[2] ^= y[0] &^ y[4]
	x[3] ^= y[1] &^ y[0]
	x[4] ^= y[2] &^ y[1]
}

func rt(x, r uint64) uint64 {
	return x<<r ^ x>>(64-r)
}

func rx1(x *[5]uint64) { // rotate & xor
	y := *x
	x[0] = y[1] ^ y[4] ^ rt(y[0], 19)
	x[1] = y[2] ^ y[0] ^ rt(y[1], 19)
	x[2] = y[3] ^ y[1] ^ rt(y[2], 19)
	x[3] = y[4] ^ y[2] ^ rt(y[3], 19)
	x[4] = y[0] ^ y[3] ^ rt(y[4], 19)
}

func rx2(x *[5]uint64) {
	y := *x
	x[0] = y[3] ^ y[2] ^ rt(y[0], 21)
	x[1] = y[4] ^ y[3] ^ rt(y[1], 21)
	x[2] = y[0] ^ y[4] ^ rt(y[2], 21)
	x[3] = y[1] ^ y[0] ^ rt(y[3], 21)
	x[4] = y[2] ^ y[1] ^ rt(y[4], 21)
}

const c13 = 1<<52 + 1<<39 + 1<<26 + 1<<13 + 1

type Bob struct {
	x [5]uint64
}

func (b *Bob) perm() {
	x := &b.x
	chi1(x)
	x[4] -= c13
	rx2(x)
	x[3] += c13
	chi2(x)
	x[2] -= c13
	rx1(x)
}

func (b *Bob) Reset() {
	for i := 0; i < 5; i++ {
		b.x[i] = 0
	}
}

func (b *Bob) Absorb(c, d uint64) {
	b.x[0] -= c
	b.perm()
	b.x[1] += d
	b.perm()
}

func (b *Bob) Squeez() (c, d uint64) {
	c = b.x[0]
	b.perm()
	d = b.x[1]
	b.perm()
	return
}
