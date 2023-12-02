package grand

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

// Random is a wrapper around a math/rand random generator
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

// Fill generates random data to fill a byte array to its length. It works
// similarly to Read([]byte) from math/rand, except it's thread safe and it
// does not return an error to be handled
func (r *Random) Fill(dest []byte) []byte {
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
	return dest
}

// ########################################################################## //
// ############################### math/rand ################################ //
// ########################################################################## //

// ExpFloat64 calls ExpFloat64 from math/rand
func (r *Random) ExpFloat64() float64 {
	return r.rng.ExpFloat64()
}

// Float32 calls Float32 from math/rand
func (r *Random) Float32() float32 {
	return r.rng.Float32()
}

// Float64 calls Float64 from math/rand
func (r *Random) Float64() float64 {
	return r.rng.Float64()
}

// Int calls Int from math/rand
func (r *Random) Int() int {
	return r.rng.Int()
}

// Intn calls Intn from math/rand
func (r *Random) Intn(n int) int {
	return r.rng.Intn(n)
}

// Int31 calls Int31 from math/rand
func (r *Random) Int31() int32 {
	return r.rng.Int31()
}

// Int31n calls Int31n from math/rand
func (r *Random) Int31n(n int32) int32 {
	return r.rng.Int31n(n)
}

// Int63 calls Int63 from math/rand
func (r *Random) Int63() int64 {
	return r.rng.Int63()
}

// Int63n calls Int63n from math/rand
func (r *Random) Int63n(n int64) int64 {
	return r.rng.Int63n(n)
}

// NormFloat64 calls NormFloat64 from math/rand
func (r *Random) NormFloat64() float64 {
	return r.rng.NormFloat64()
}

// Perm calls Perm from math/rand
func (r *Random) Perm(n int) []int {
	return r.rng.Perm(n)
}

// Read calls Read from math/rand
func (r *Random) Read(p []byte) (n int, err error) {
	return r.rng.Read(p)
}

// Seed calls Seed from math/rand
func (r *Random) Seed(seed int64) {
	r.rng.Seed(seed)
}

// Shuffle calls Shuffle from math/rand
func (r *Random) Shuffle(n int, swap func(i, j int)) {
	r.rng.Shuffle(n, swap)
}

// Uint32 calls Uint32 from math/rand
func (r *Random) Uint32() uint32 {
	return r.rng.Uint32()
}

// Uint64 calls Uint64 from math/rand
func (r *Random) Uint64() uint64 {
	return r.rng.Uint64()
}
