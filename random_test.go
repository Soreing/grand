package grand

import (
	"fmt"
	"io"
	"testing"

	crand "crypto/rand"
	"math/rand"

	"github.com/stretchr/testify/assert"
)

type errReader struct{}

func (r *errReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

// TestNewSource tests that an source is created without seed
func TestNewSource(t *testing.T) {
	tests := []struct {
		Name   string
		Reader io.Reader
		Error  error
	}{
		{
			Name:   "New Source without given seed",
			Reader: crand.Reader,
			Error:  nil,
		},
		{
			Name:   "New Source without given seed and error",
			Reader: &errReader{},
			Error:  fmt.Errorf("failed to create seed"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			crand.Reader = test.Reader
			src, err := NewSource()

			assert.Equal(t, test.Error, err)
			if err == nil {
				assert.NotNil(t, src)
			}
		})
	}
}

// TestNewSource tests that an source is created with seed
func TestNewSourceWithSeed(t *testing.T) {
	tests := []struct {
		Name   string
		Seed   int64
		Number int64
		Error  error
	}{
		{
			Name:   "New Source without given seed",
			Seed:   1275028672939391351,
			Number: 5129775219661360826,
			Error:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			src, err := NewSource(int64(test.Seed))

			assert.NotNil(t, src)
			assert.Equal(t, test.Error, err)
			assert.Equal(t, test.Number, src.Int63())
		})
	}
}

// TestNewRandom tests that a randomizer is created from source
func TestNewRandom(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{
			Name: "New Random",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			src, _ := NewSource(0)
			rand := New(src)

			assert.NotNil(t, rand)
			assert.NotNil(t, rand.rng)
		})
	}
}

// TestFill tests that Fill generates random bytes and fills a byte array
func TestFill(t *testing.T) {
	tests := []struct {
		Name   string
		Seed   int64
		Array  []byte
		Result []byte
	}{
		{
			Name:  "Fill byte array",
			Seed:  1275028672939391351,
			Array: make([]byte, 10),
			Result: []byte{
				0xba, 0x4e, 0x84, 0x31, 0x62,
				0x9f, 0x30, 0x76, 0xf3, 0x78,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			src, _ := NewSource(test.Seed)
			rand := New(src)

			res := rand.Fill(test.Array)

			assert.Equal(t, test.Result, res)
		})
	}
}

// TestExpFloat64 tests that the result of ExpFloat64 matches with math/rand
func TestExpFloat64(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.ExpFloat64()
				act := mrng.ExpFloat64()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestFloat32 tests that the result of Float32 matches with math/rand
func TestFloat32(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Float32()
				act := mrng.Float32()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestFloat64 tests that the result of Float64 matches with math/rand
func TestFloat64(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Float64()
				act := mrng.Float64()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestInt tests that the result of Int matches with math/rand
func TestInt(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Int()
				act := mrng.Int()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestIntn tests that the result of Intn matches with math/rand
func TestIntn(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
		Max  int
	}{
		{
			Name: "Default Seed",
			Seed: 0,
			Max:  100000000,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
			Max:  123456789,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Intn(test.Max)
				act := mrng.Intn(test.Max)

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestInt31 tests that the result of Int31 matches with math/rand
func TestInt31(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Int31()
				act := mrng.Int31()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestInt31n tests that the result of Int31n matches with math/rand
func TestInt31n(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
		Max  int32
	}{
		{
			Name: "Default Seed",
			Seed: 0,
			Max:  100000000,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
			Max:  123456789,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Int31n(test.Max)
				act := mrng.Int31n(test.Max)

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestInt63 tests that the result of Int63 matches with math/rand
func TestInt63(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Int63()
				act := mrng.Int63()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestInt63n tests that the result of Int63n matches with math/rand
func TestInt63n(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
		Max  int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
			Max:  100000000,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
			Max:  123456789,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Int63n(test.Max)
				act := mrng.Int63n(test.Max)

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestNormFloat64 tests that the result of NormFloat64 matches with math/rand
func TestNormFloat64(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.NormFloat64()
				act := mrng.NormFloat64()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestPerm tests that the result of Perm matches with math/rand
func TestPerm(t *testing.T) {
	tests := []struct {
		Name  string
		Seed  int64
		Count int
	}{
		{
			Name:  "Default Seed",
			Seed:  0,
			Count: 10,
		},
		{
			Name:  "Custom Seed",
			Seed:  1275028672939391351,
			Count: 15,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Perm(test.Count)
				act := mrng.Perm(test.Count)

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestRead tests that the result of Read matches with math/rand
func TestRead(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
		Size int
	}{
		{
			Name: "Default Seed",
			Seed: 0,
			Size: 10,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
			Size: 15,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp, en := grng.Read(make([]byte, test.Size))
				act, an := mrng.Read(make([]byte, test.Size))

				assert.Equal(t, en, an)
				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestSeed tests that the result of Seed matches with math/rand
func TestSeed(t *testing.T) {
	tests := []struct {
		Name    string
		Seed    int64
		NewSeed int
	}{
		{
			Name:    "Default Seed",
			Seed:    0,
			NewSeed: 7127123515134123452,
		},
		{
			Name:    "Custom Seed",
			Seed:    1275028672939391351,
			NewSeed: 1092375102331312441,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				grng.Seed(int64(test.NewSeed))
				exp := grng.Int()
				mrng.Seed(int64(test.NewSeed))
				act := mrng.Int()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestShuffle tests that the result of Shuffle matches with math/rand
func TestShuffle(t *testing.T) {
	tests := []struct {
		Name  string
		Seed  int64
		Input []byte
	}{
		{
			Name: "Default Seed",
			Seed: 0,
			Input: []byte{
				0xba, 0x4e, 0x84, 0x31, 0x62,
				0x9f, 0x30, 0x76, 0xf3, 0x78,
			},
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
			Input: []byte{
				0xba, 0x4e, 0x84, 0x31, 0x62,
				0x9f, 0x30, 0x76, 0xf3, 0x78,
				0xba, 0x4e, 0x84, 0x31, 0x62,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)
			gslc := make([]byte, len(test.Input))
			copy(gslc, test.Input)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)
			mslc := make([]byte, len(test.Input))
			copy(mslc, test.Input)

			for i := 0; i < 100; i++ {
				grng.Shuffle(len(gslc), func(i, j int) {
					tmp := gslc[i]
					gslc[i] = gslc[j]
					gslc[j] = tmp
				})

				mrng.Shuffle(len(gslc), func(i, j int) {
					tmp := mslc[i]
					mslc[i] = mslc[j]
					mslc[j] = tmp
				})

				assert.Equal(t, gslc, mslc)
			}
		})
	}
}

// TestUint32 tests that the result of Uint32 matches with math/rand
func TestUint32(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Uint32()
				act := mrng.Uint32()

				assert.Equal(t, exp, act)
			}
		})
	}
}

// TestUint64 tests that the result of Uint64 matches with math/rand
func TestUint64(t *testing.T) {
	tests := []struct {
		Name string
		Seed int64
	}{
		{
			Name: "Default Seed",
			Seed: 0,
		},
		{
			Name: "Custom Seed",
			Seed: 1275028672939391351,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			gsrc, _ := NewSource(test.Seed)
			grng := New(gsrc)

			msrc := rand.NewSource(test.Seed)
			mrng := rand.New(msrc)

			for i := 0; i < 100; i++ {
				exp := grng.Uint64()
				act := mrng.Uint64()

				assert.Equal(t, exp, act)
			}
		})
	}
}
