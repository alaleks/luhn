package luhn_test

import (
	"testing"

	"github.com/alaleks/luhn"
)

const num = 10000

func TestValid(t *testing.T) {
	luhnNumber := 79927398713
	check := luhn.Valid(luhnNumber)

	if !check {
		t.Errorf("%d must match Luhn's algorithm", luhnNumber)
	}
}

func TestGenerateEven(t *testing.T) {
	for i := 0; i < num; i++ {
		if l := luhn.Generate(12); !luhn.Valid(l) {
			t.Errorf("a number was received that does not match the Luhn algorithm: %d", l)
		}
	}
}

func TestGenerateOdd(t *testing.T) {
	for i := 0; i < num; i++ {
		if l := luhn.Generate(11); !luhn.Valid(l) {
			t.Errorf("a number was received that does not match the Luhn algorithm: %d", l)
		}
	}
}

func BenchmarkGenerateLuhn8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		luhn.Generate(8)
	}
}

func BenchmarkGenerateLuhn12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		luhn.Generate(12)
	}
}

func BenchmarkGenerateLuhn16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		luhn.Generate(16)
	}
}
