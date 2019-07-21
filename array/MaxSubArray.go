package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	var low, high, sum = maxSubArray(a, 0, len(a)-1)
	fmt.Println("Position = [" + strconv.Itoa(low) + ";" + strconv.Itoa(high) + "] sum = " + strconv.Itoa(sum))
	low, high, sum = maxSubArrayBrutalForce(a)
	fmt.Println("Position = [" + strconv.Itoa(low) + ";" + strconv.Itoa(high) + "] sum = " + strconv.Itoa(sum))
	low, high, sum = maxSubArrayDynamicPrograming(a)
	fmt.Println("Position = [" + strconv.Itoa(low) + ";" + strconv.Itoa(high) + "] sum = " + strconv.Itoa(sum))
}

func maxSubArrayBrutalForce(a []int) (int, int, int) {
	var max = a[0]
	var fromIndex = 0
	var toIndex = 0
	for i := 0; i < len(a)-2; i++ {
		tempMax := a[i]
		for j := i + 1; j < len(a)-1; j++ {
			tempMax = tempMax + a[j]
			if max < tempMax {
				max = tempMax
				fromIndex = i
				toIndex = j
			}
		}
	}
	return fromIndex, toIndex, max
}
func maxSubArray(a []int, low int, high int) (int, int, int) {
	if low == high {
		return low, high, a[low]
	}
	// Find Mid first
	mid := (low + high) / 2
	// We find LeftMaxSubArray
	leftLow, leftHigh, sumLeft := maxSubArray(a, low, mid)
	// We find RightMaxSubArray
	rightLow, rightHigh, rightSum := maxSubArray(a, mid+1, high)
	// We find maxCrossingSubArray
	crossLow, crossHigh, crossSum := maxCrossingSubArray(a, leftLow, mid, rightHigh)
	// We find Max of [leftMaxSubArray,RightMaxSubArray,MaxCrossSubArray] and return result
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
	var lowSum, sum, highSum, crossLow, crossHigh int
	sum = 0
	// we find max SubArray end with mid and start from Low
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > lowSum || i == mid {
			lowSum = sum
			crossLow = i
		}
	}
	sum = 0
	// we find max SubArray start with mid and end with High
	for i := (mid + 1); i <= high; i++ {
		sum = sum + a[i]
		if sum > highSum || i == mid+1 {
			highSum = sum
			crossHigh = i
		}
	}
	return crossLow, crossHigh, lowSum + highSum
}

type MaxSub struct {
	low, high, sum int
}

func maxSubArrayDynamicProgramingRecursive(a []int, cache map[int]MaxSub) {

}
func sumArrayByIndex(a []int, cache map[int]MaxSub, index int) MaxSub {
	if index == 0 {
		var maxSub = MaxSub{0, 0, a[0]}
		cache[0] = maxSub
		return maxSub
	}
	var previousSumArray = sumArrayByIndex(a, cache, index-1)
	//var newSumArray = new MaxSub{previousSumArray.}
}

func maxSubArrayDynamicPrograming(a []int) (int, int, int) {
	var max, lowMax, highMax, sum, lowSum, highSum int
	for i := 0; i < len(a); i++ {
		if i == 0 {
			max = a[i]
			lowMax = i
			highMax = i
			sum = a[i]
			lowSum = i
			highSum = i
		} else {
			sum = a[i] + sum
			if a[i] > sum {
				if a[i] >= max {
					max = a[i]
					lowMax = i
					highMax = i
				}
				sum = a[i]
				lowSum = i
				highSum = i
			} else {
				highSum = i
				if sum >= max {
					max = sum
					highMax = highSum
					lowMax = lowSum
				}
			}
		}
	}
	return lowMax, highMax, max
}
