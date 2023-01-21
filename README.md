# luhn

**luhn** is a simple implementation of Luhn's algorithm: number generation and validation.

## Installation

```
go get github.com/alaleks/luhn
```

## Example

```
package main

import (
	"fmt"

	"github.com/alaleks/luhn"
)

func main() {
	luhnNumberEven := luhn.Generate(12)
	luhnNumberOdd := luhn.Generate(17)

	fmt.Printf("%d: %v\n", luhnNumberEven, luhn.Valid(luhnNumberEven))
	fmt.Printf("%d: %v\n", luhnNumberOdd, luhn.Valid(luhnNumberOdd))
}
```

## Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/alaleks/luhn
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
BenchmarkGenerateLuhn8-12          21308             56077 ns/op
BenchmarkGenerateLuhn12-12         13784             84676 ns/op
BenchmarkGenerateLuhn16-12         10420            115176 ns/op
```
