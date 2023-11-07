package grand

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// Random is a wrapper around a math/rand random generator initialized by
// providing a seed or reading data from an entropy source.
type Random struct {
	rng *rand.Rand
}

// NewSource creates a random source. If a seed is given, it is used to seed the
// random generator, otherwise entropy data from a crypto/rand is used as seed.
func NewSource(seed ...int64) (rand.Source, error) {
	if len(seed) > 0 {
		return rand.NewSource(seed[0]), nil
	} else {
		entropy := make([]byte, 8)
		n, err := crand.Read(entropy)
		if err != nil || n != 8 {
			return nil, fmt.Errorf("failed to create seed")
		}
		s := int64(binary.BigEndian.Uint64(entropy))
		return rand.NewSource(s), nil
	}
}

// New creates a random generator from a source.
func New(src rand.Source) *Random {
	return &Random{rand.New(src)}
}

// Fill generates random data to fill a byte array to its length.
func (r *Random) Fill(dest []byte) {
	bytes, value := 0, int64(0)
	for i := 0; i < len(dest); i++ {
		if bytes == 0 {
			value = r.rng.Int63()
			bytes = 7
		}
		dest[i] = byte(value & 0xFF)
		value >>= 8
		bytes--
	}
}
