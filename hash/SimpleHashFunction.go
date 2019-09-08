package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("****testHashByDivision**")
	testHashByDivision()
	fmt.Println("****testByMultiplication**")
	testByMultiplication()
}
func testHashByDivision() {
	tablesize := 15
	var primeNumber = findNearestPrime(tablesize)
	fmt.Println(hashByDivisionMethod("test", primeNumber))
	fmt.Println(hashByDivisionMethod("test2", primeNumber))
	fmt.Println(hashByDivisionMethod("1test", primeNumber))
	fmt.Println(hashByDivisionMethod("tafest", primeNumber))
	fmt.Println(hashByDivisionMethod("te4st", primeNumber))
	fmt.Println(hashByDivisionMethod("tesfst", primeNumber))
	fmt.Println(hashByDivisionMethod("raaf", primeNumber))
	fmt.Println(hashByDivisionMethod("tesqwrfst", primeNumber))
	fmt.Println(hashByDivisionMethod("tesfst", primeNumber))
	fmt.Println(hashByDivisionMethod("tesfsast", primeNumber))
	fmt.Println(hashByDivisionMethod("tesaafst", primeNumber))
	fmt.Println(hashByDivisionMethod("te124sfst", primeNumber))
	fmt.Println(hashByDivisionMethod("tsaesfst", primeNumber))
}
func testByMultiplication() {
	tablesize := 15
	fmt.Println(hashByMultiplicationMethod("test", tablesize))
	fmt.Println(hashByMultiplicationMethod("test2", tablesize))
	fmt.Println(hashByMultiplicationMethod("1test", tablesize))
	fmt.Println(hashByMultiplicationMethod("tafest", tablesize))
	fmt.Println(hashByMultiplicationMethod("te4st", tablesize))
	fmt.Println(hashByMultiplicationMethod("tesfst", tablesize))
	fmt.Println(hashByMultiplicationMethod("raaf", tablesize))
	fmt.Println(hashByMultiplicationMethod("tesqwrfst", tablesize))
	fmt.Println(hashByMultiplicationMethod("tesfst", tablesize))
	fmt.Println(hashByMultiplicationMethod("tesfsast", tablesize))
	fmt.Println(hashByMultiplicationMethod("tesaafst", tablesize))
	fmt.Println(hashByMultiplicationMethod("te124sfst", tablesize))
	fmt.Println(hashByMultiplicationMethod("tsaesfst", tablesize))
}
func hashByDivisionMethod(key string, primeNumber int) int32 {
	keyChars := []rune(key)
	var keyChar int32
	var hash int32
	for i := 0; i < len(keyChars); i++ {
		keyChar = keyChars[i]
		hash = (hash + keyChar) % int32(primeNumber)
	}
	return hash
}

func hashByMultiplicationMethod(key string, tableSize int) int32 {
	keyChars := []rune(key)
	var keyChar int32
	constantA := 0.618
	for i := 0; i < len(keyChars); i++ {
		// try to make key become different
		keyChar = keyChar + keyChars[i] + keyChar%11
	}
	hashResult := float64(tableSize) * math.Mod(float64(keyChar)*constantA, 1)
	return int32(math.Floor(hashResult))
}

func findNearestPrime(number int) int {
	primeNumber := 1
	for i := 0; i < number; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		primeNumber = i
	}
	return primeNumber
}
