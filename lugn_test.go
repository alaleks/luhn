package luhn_test

import (
	"testing"

	"github.com/alaleks/luhn"
)

func TestValid(t *testing.T) {
	luhnNumber := 79927398713
	check := luhn.Valid(luhnNumber)

	if !check {
		t.Errorf("%d must match Luhn's algorithm", luhnNumber)
	}
}

func TestGenerateEven(t *testing.T) {
	for i := 1; i < 10001; i++ {
		if l := luhn.Generate(12); !luhn.Valid(l) {
			t.Errorf("a number was received that does not match the Luhn algorithm: %d", l)
		}
	}
}

func TestGenerateOdd(t *testing.T) {
	for i := 1; i < 10001; i++ {
		if l := luhn.Generate(11); !luhn.Valid(l) {
			t.Errorf("a number was received that does not match the Luhn algorithm: %d", l)
		}
	}
}
