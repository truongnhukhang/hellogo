package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(random(5, 9))
	fmt.Println(random(6, 10))
}

func random(a int, b int) int {
	if a == b {
		return a
	}
	mid := (b + a) / 2
	rand01 := random01()
	fmt.Println("mid : " + strconv.Itoa(mid) + " - rand01 : " + strconv.Itoa(rand01))
	if rand01 == 0 {
		return random(a, mid)
	} else {
		return random(mid+1, b)
	}
}

func random01() int {
	return rand.Intn(2)
}
