package luhn

import (
	"math/rand"
	"time"
)

// An integer is passed to the function,
// which indicates the length of the number corresponding
// to the Luhn Algorithm that will be returned.
func Generate(n int) int {
	var (
		numbers = make([]int, n)
		res     int
		luhn    int
		count   int
	)

	if n%2 != 0 {
		count = 1
	}

	for i := range numbers {
		rand.Seed(time.Now().UnixNano())
		numbers[i] = rand.Intn(9) + 1

		if count%2 != 0 {
			luhn += numbers[i]
			count++
			continue
		}

		number := numbers[i] * 2

		if number > 9 {
			number -= 9
		}

		luhn += number
		count++
	}

	if luhn%10 <= numbers[n-1] {
		numbers[n-1] -= luhn % 10
	} else {
		numbers[n-1] += (10 - luhn%10)
	}

	for _, v := range numbers {
		res = 10*res + v
	}

	return res
}

// Checking a Number for Compliance with the Luhn Algorithm
func Valid(n int) bool {
	var (
		luhn       = 0
		number     = n / 10
		checkDigit = n % 10
	)

	for i := 0; number > 0; i++ {
		d := number % 10

		if i%2 == 0 {
			d = d * 2
			if d > 9 {
				d -= 9
			}
		}

		luhn += d
		number = number / 10
	}

	return (checkDigit+luhn)%10 == 0
}
