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
	var mid = (high + low) / 2
	var leftLow, leftHigh, leftSum = maxSubArray(a, low, mid)
	var rightLow, rightHigh, rightSum = maxSubArray(a, mid+1, high)
	var crossLow, crossHigh, crossSum = maxCrossingSubArray(a, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}

func maxCrossingSubArray(a []int, low int, mid int, high int) (int, int, int) {
	var leftSum = 0
	var sum = 0
	var maxLeft = low
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	var rightSum = 0
	sum = 0
	var maxRight = high
	for j := mid + 1; j <= high; j++ {
		sum = sum + a[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}
