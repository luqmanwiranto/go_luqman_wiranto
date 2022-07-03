package main

import "fmt"

var memoTopDown = map[int]int{}

func fiboTopDown(n int) (fiboNumber int) {
	valueFibo, isAlreadyCount := memoTopDown[n]
	if isAlreadyCount {
		return valueFibo
	}
	if n <= 1 {
		memoTopDown[n] = n
	} else {
		memoTopDown[n] = fiboTopDown(n-1) + fiboTopDown(n-2)
	}
	return memoTopDown[n]
}

func main() {
	fmt.Println(fiboTopDown(5))
}
