package main

import "fmt"

func binarySearch(sortArr []int, x int) (index int) {
	iKiri, iKanan := 0, len(sortArr)-1
	for iKiri <= iKanan {
		iTengah := (iKiri + iKanan) / 2
		if x < sortArr[iTengah] {
			iKanan = iTengah - 1
		} else if x > sortArr[iTengah] {
			iKiri = iTengah + 1
		} else if x == sortArr[iTengah] {
			return iTengah
		}
	}
	return index
}

func main() {
	fmt.Println(binarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Println(binarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100))
}
