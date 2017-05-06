/*	Bob is a little sponge ;)

	Author: Serhat Sevki Dincer jfcgaussATgmail
*/

package bob

func rt(x, r uint64) uint64 {
	return x<<r ^ x>>(64-r)
}

func chi1(x *[7]uint64) { // non-linear
	y := *x
	x[0] = y[1]&^y[6] ^ rt(y[0], 7)
	x[1] = y[2]&^y[0] ^ rt(y[1], 7)
	x[2] = y[3]&^y[1] ^ rt(y[2], 7)
	x[3] = y[4]&^y[2] ^ rt(y[3], 7)
	x[4] = y[5]&^y[3] ^ rt(y[4], 7)
	x[5] = y[6]&^y[4] ^ rt(y[5], 7)
	x[6] = y[0]&^y[5] ^ rt(y[6], 7)
}

func chi2(x *[7]uint64) {
	y := *x
	x[0] = y[5]&^y[2] ^ rt(y[0], 19)
	x[1] = y[6]&^y[3] ^ rt(y[1], 19)
	x[2] = y[0]&^y[4] ^ rt(y[2], 19)
	x[3] = y[1]&^y[5] ^ rt(y[3], 19)
	x[4] = y[2]&^y[6] ^ rt(y[4], 19)
	x[5] = y[3]&^y[0] ^ rt(y[5], 19)
	x[6] = y[4]&^y[1] ^ rt(y[6], 19)
}

func chi3(x *[7]uint64) {
	y := *x
	x[0] = y[3]&^y[4] ^ rt(y[0], 33)
	x[1] = y[4]&^y[5] ^ rt(y[1], 33)
	x[2] = y[5]&^y[6] ^ rt(y[2], 33)
	x[3] = y[6]&^y[0] ^ rt(y[3], 33)
	x[4] = y[0]&^y[1] ^ rt(y[4], 33)
	x[5] = y[1]&^y[2] ^ rt(y[5], 33)
	x[6] = y[2]&^y[3] ^ rt(y[6], 33)
}

func rx1(x *[7]uint64) { // rotate & xor
	y := *x
	x[0] = y[1] ^ y[6] ^ rt(y[0], 57)
	x[1] = y[2] ^ y[0] ^ rt(y[1], 57)
	x[2] = y[3] ^ y[1] ^ rt(y[2], 57)
	x[3] = y[4] ^ y[2] ^ rt(y[3], 57)
	x[4] = y[5] ^ y[3] ^ rt(y[4], 57)
	x[5] = y[6] ^ y[4] ^ rt(y[5], 57)
	x[6] = y[0] ^ y[5] ^ rt(y[6], 57)
}

func rx2(x *[7]uint64) {
	y := *x
	x[0] = y[5] ^ y[2] ^ rt(y[0], 45)
	x[1] = y[6] ^ y[3] ^ rt(y[1], 45)
	x[2] = y[0] ^ y[4] ^ rt(y[2], 45)
	x[3] = y[1] ^ y[5] ^ rt(y[3], 45)
	x[4] = y[2] ^ y[6] ^ rt(y[4], 45)
	x[5] = y[3] ^ y[0] ^ rt(y[5], 45)
	x[6] = y[4] ^ y[1] ^ rt(y[6], 45)
}

func rx3(x *[7]uint64) {
	y := *x
	x[0] = y[3] ^ y[4] ^ rt(y[0], 31)
	x[1] = y[4] ^ y[5] ^ rt(y[1], 31)
	x[2] = y[5] ^ y[6] ^ rt(y[2], 31)
	x[3] = y[6] ^ y[0] ^ rt(y[3], 31)
	x[4] = y[0] ^ y[1] ^ rt(y[4], 31)
	x[5] = y[1] ^ y[2] ^ rt(y[5], 31)
	x[6] = y[2] ^ y[3] ^ rt(y[6], 31)
}

const c13 = 1<<52 + 1<<39 + 1<<26 + 1<<13 + 1

type Bob struct {
	x [7]uint64
}

func (b *Bob) perm() {
	x := &b.x
	chi1(x)
	x[2] -= c13
	rx2(x)
	x[4] += c13
	chi3(x)
	x[6] -= c13
	rx1(x)
	x[1] += c13
	chi2(x)
	x[3] -= c13
	rx3(x)
	x[5] += c13
}

func (b *Bob) Reset() {
	for i := 6; i >= 0; i-- {
		b.x[i] = 0
	}
}

func (b *Bob) Absorb(i uint64) {
	b.x[0] += i
	b.perm()
}

func (b *Bob) Squeez() uint64 {
	o := b.x[0]
	b.perm()
	return o
}
