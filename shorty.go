package shorty

import (
	"math/rand"
	"time"
)

// plan B with charCodes: 65 -> 90 A-Z, 97 -> 122 a-z
var charpool = []byte("abcdefghijklmnopqrstuvxyz1234567890")

// randByte returns a random char from charpool
func randByte() byte {
	return charpool[rand.Intn(len(charpool))]
}

// randSeq returns n random bytes from charpool
func randSeq(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = randByte()
	}
	return out
}
