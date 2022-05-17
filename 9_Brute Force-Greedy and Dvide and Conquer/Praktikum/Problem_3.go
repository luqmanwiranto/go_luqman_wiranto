package main

import (
	"fmt"
	"sort"
)

func main() {
	var dragonHead, knightHeight []int
	dragonHead = []int{5, 4}
	knightHeight = []int{7, 8, 4}
	DragonOfLoowater(dragonHead, knightHeight)
}

func DragonOfLoowater(dragonHead, knightHeight []int) {
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("kningt fall")
	}
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	totalMinimumHeightOfKnight := 0
	for i := 0; i < len(dragonHead); i++ {
		for j := 0; j < len(knightHeight); j++ {
			if knightHeight[j] >= dragonHead[i] {
				totalMinimumHeightOfKnight += knightHeight[j]
				break
			}
		}
	}
	if totalMinimumHeightOfKnight == 0 {
		fmt.Println("knight fall")
	} else {
		fmt.Println(totalMinimumHeightOfKnight)
	}
	return
}
