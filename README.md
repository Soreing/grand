# Go Random Generator

![Build](https://github.com/soreing/grand/actions/workflows/build_status.yaml/badge.svg)
![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/Soreing/4b6f950f01f3e6e5b9ed17b268664538/raw/grand)
[![Go Report Card](https://goreportcard.com/badge/github.com/Soreing/grand)](https://goreportcard.com/report/github.com/Soreing/grand)
[![Go Reference](https://pkg.go.dev/badge/github.com/Soreing/grand.svg)](https://pkg.go.dev/github.com/Soreing/grand)

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
