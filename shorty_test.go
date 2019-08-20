package shorty

import "testing"

func TestRandSeq(t *testing.T) {
	n := 5
	seq := randSeq(n)
	// should be of n length
	if len(seq) != n {
		t.Errorf("expected equality between %d and %d", len(seq), n)
	}
	// should only include allowed chars
	for _, b := range seq {
		if !allowed(b) {
			t.Errorf("%v is not allowed", b)
		}
	}
}

func allowed(char uint8) bool {
	for i := 0; i < len(charpool); i++ {
		if char == charpool[i] {
			return true
		}
	}
	return false
}
