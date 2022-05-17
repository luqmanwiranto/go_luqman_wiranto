package main

import (
	"fmt"
	"sort"
)

func main() {
	MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})
}

func MaxSequence(arr []int) int {
	sum := []int{}
	hmap := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		arrGroup := []int{arr[i]}
		tmpSum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			arrGroup = append(arrGroup, arr[j])
			tmpSum += arr[j]
		}
		sum = append(sum, tmpSum)
		hmap[tmpSum] = arrGroup
		tempArr := arr
		for len(tempArr) != i {
			arrGroup = []int{arr[i]}
			tmpSum += tempArr[i]
		}
		sum = append(sum, tmpSum)
		hmap[tmpSum] = arrGroup
		tempArr = tempArr[:len(tempArr)-1]
	}
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	fmt.Println(sum[0])
	fmt.Println(hmap[sum[0]])
	return sum[0]
}
