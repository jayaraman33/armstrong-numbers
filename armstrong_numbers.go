package armstrong

import "math"

func IsNumber(n int) bool {

	var len, sum int

	for num := n; num > 0; {
		num = num / 10
		len++
	}

	for num := n; num > 0; {
		digit := num % 10
		num = num / 10

		sum += int(math.Pow(float64(digit), float64(len)))
	}
	return sum == n
}
