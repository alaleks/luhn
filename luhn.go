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
		op      = 1
	)

	if n%2 == 0 {
		for i := range numbers {
			rand.Seed(time.Now().UnixNano())
			numbers[i] = rand.Intn(9) + 1

			if (i+1)%2 == 0 {
				luhn += numbers[i]
				continue
			}

			number := numbers[i] * 2

			if number > 9 {
				number -= 9
			}

			luhn += number
		}
	} else {
		for i := range numbers {
			rand.Seed(time.Now().UnixNano())
			numbers[i] = rand.Intn(9) + 1

			if (i+1)%2 != 0 {
				luhn += numbers[i]
				continue
			}

			number := numbers[i] * 2

			if number > 9 {
				number -= 9
			}

			luhn += number
		}
	}

	if luhn%10 <= numbers[len(numbers)-1] {
		numbers[len(numbers)-1] -= luhn % 10
	} else {
		numbers[len(numbers)-1] += (10 - luhn%10)
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		res += numbers[i] * op
		op *= 10
	}

	return res
}

// Checking a Number for Compliance with the Luhn Algorithm
func Valid(n int) bool {
	var (
		luhSum     = 0
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

		luhSum += d
		number = number / 10
	}

	return (checkDigit+luhSum)%10 == 0
}
