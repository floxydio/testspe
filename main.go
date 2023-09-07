package main

import (
	"testspe/blueocean"
	"testspe/findneedle"
	"testspe/narcisstic"
	"testspe/oddeven"
)

func main() {
	narcisstic.NarcisticCount(153)
	narcisstic.NarcisticCount(111)

	oddeven.OddEvenCheck([]int{2, 4, 0, 100, 11, 2602, 36})
	oddeven.OddEvenCheck([]int{160, 3, 1719, 19, 11, 13, -21})

	findneedle.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue")
	blueocean.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1})
	blueocean.BlueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5})

}
