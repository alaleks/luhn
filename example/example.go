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
