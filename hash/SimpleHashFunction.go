package main

import "fmt"

func main() {
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
