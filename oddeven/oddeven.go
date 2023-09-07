package oddeven

func OddEvenCheck(number []int) int {
	evenCount := 0
	oddCount := 0
	lastEven := 0
	lastOdd := 0

	for _, num := range number {
		if num%2 == 0 {
			evenCount++
			lastEven = num
		} else {
			oddCount++
			lastOdd = num
		}
		if evenCount > 1 && oddCount == 1 {
			return lastOdd
		} else if oddCount > 1 && evenCount == 1 {
			return lastEven
		}
	}
	return 0

}
