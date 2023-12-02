# Go Random Generator

Grand is a simple random generator wrapper on math/rand seeded by entropy from crypto/rand

## Usage

Create a source that can be used by the randomizer. `NewSource` creates a source
with an optional seed. If there is no seed provided, a truly random seed will be 
derived from crypto/rand.

If a seed is provided, the function will not return an error. If entropy is used,
the function can return an error if it fails to read bytes from crypto/rand

```golang
// New source using entropy
src, err := grand.NewSource()

// New source from existing seed
src, _ := grand.NewSource(1552664635341843643)
```

The randomizer can use any math/rand source, not only ones created using this
package.

```golang
// New randomizer from source
rng := grand.New(src)
```

The Fill method fills a destination byte array with random values like `Read`
from math/rand, but unlike Read, it is thread safe and can be called concurrently.
It also does not return an error to be handled.

```golang
// Fill slice with random data
dst := make([]byte, 20)
rng.Fill(dst)

// Prints 20 random bytes
fmt.Println(dst)
```
