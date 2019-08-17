package shorty

import (
  "math/rand"
  "time"
)

// charCodes: 65 -> 90 A-Z, 97 -> 122 a-z
var charpool = []byte("abcdefghijklmnopqrstuvxyz1234567890")

func randChar() byte {
	return charpool[rand.Intn(len(charpool))]
}

func RandSeq(n int) []byte {
  rand.Seed(time.Now().UnixNano())

  var out []byte
  for i := 0; i < n; i++ {
    out = append(out, randChar())
  }
  return out
}
