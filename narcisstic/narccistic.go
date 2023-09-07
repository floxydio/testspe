package narcisstic

import "math"

func NarcisticCount(number int) bool {
	numberDigit := CountNumber(number)
	sum := 0
	originalNumber := number

	for number > 0 {
		digit := number % 10
		sum += int(math.Pow(float64(digit), float64(numberDigit)))
		number /= 10
	}

	return sum == originalNumber

}

func CountNumber(number int) int {
	count := 0

	for number > 0 {
		number /= 10
		count++
	}
	return count
}
