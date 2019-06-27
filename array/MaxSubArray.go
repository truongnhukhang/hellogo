package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	var low, high, sum = maxSubArray(a, 0, len(a)-1)
	fmt.Println("Position = [" + strconv.Itoa(low) + ";" + strconv.Itoa(high) + "] sum = " + strconv.Itoa(sum))
}

func maxSubArray(a []int, low int, high int) (int, int, int) {
	if low == high {
		return low, high, a[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, sumLeft := maxSubArray(a, low, mid)
	rightLow, rightHigh, rightSum := maxSubArray(a, mid+1, high)
	crossLow, crossHigh, crossSum := maxCrossingSubArray(a, leftHigh, mid, rightLow)
	if sumLeft >= rightSum {
		if rightSum >= crossSum {
			return leftLow, leftHigh, sumLeft
		}
	} else {
		if sumLeft >= crossSum {
			return rightLow, rightHigh, rightSum
		}
	}
	return crossLow, crossHigh, crossSum
}

func maxCrossingSubArray(a []int, low int, mid int, high int) (int, int, int) {
	var lowSum, sum, highSum, crossLow, crossHigh = 0, 0, 0, 0, 0
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > lowSum || i == mid {
			lowSum = sum
			crossLow = i
		}
	}
	sum = 0
	for i := (mid + 1); i <= high; i++ {
		sum = sum + a[i]
		if sum > highSum || i == mid+1 {
			highSum = sum
			crossHigh = i
		}
	}
	return crossLow, crossHigh, (lowSum + highSum)
}
