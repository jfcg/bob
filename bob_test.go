package bob

import "testing"

func Test1(t *testing.T) {
	var b Bob
	const N = 9e7
	mp := make(map[uint64]bool, N*2)

	for i := N; i > 0; i-- {
		c, d := b.Squeez()
		mp[c] = true
		mp[d] = true
	}

	if len(mp) != N*2 {
		t.Fatal("bad Bob!", len(mp), N*2)
	}
}
