package main

import "fmt"

func main() {
	tablesize := 15
	fmt.Println(hashByDivisionMethod("test", tablesize))
	fmt.Println(hashByDivisionMethod("test2", tablesize))
	fmt.Println(hashByDivisionMethod("1test", tablesize))
	fmt.Println(hashByDivisionMethod("tafest", tablesize))
	fmt.Println(hashByDivisionMethod("te4st", tablesize))
	fmt.Println(hashByDivisionMethod("tesfst", tablesize))
	fmt.Println(hashByDivisionMethod("raaf", tablesize))
	fmt.Println(hashByDivisionMethod("tesqwrfst", tablesize))
	fmt.Println(hashByDivisionMethod("tesfst", tablesize))
	fmt.Println(hashByDivisionMethod("tesfsast", tablesize))
	fmt.Println(hashByDivisionMethod("tesaafst", tablesize))
	fmt.Println(hashByDivisionMethod("te124sfst", tablesize))
	fmt.Println(hashByDivisionMethod("tsaesfst", tablesize))
}

func hashByDivisionMethod(key string, tableSize int) int32 {
	keyChars := []rune(key)
	var keyChar int32
	var hash int32
	var primeNumber = findNearestPrime(tableSize)
	for i := 0; i < len(keyChars); i++ {
		keyChar = keyChars[i]
		hash = (31*hash + keyChar) % int32(primeNumber)
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
