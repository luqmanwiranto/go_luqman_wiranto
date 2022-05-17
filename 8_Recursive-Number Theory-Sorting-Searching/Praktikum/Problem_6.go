package main

import "fmt"

func Playingdomino(cards [][]int, deck []int) interface{} {
	idxBiggestPair := -1
	biggestPair := 0
	for i := 0; i < len(cards); i++ {
		if deck[0] == cards[i][0] && deck[0] > biggestPair ||
			deck[0] == cards[i][1] && deck[0] > biggestPair {
			biggestPair == deck[0]
			idxBiggestPair = i
		}
		if deck[0] == cards[i][0] && deck[0] > biggestPair ||
			deck[0] == cards[i][1] && deck[0] > biggestPair {
			biggestPair == deck[1]
			idxBiggestPair = i
		}
	}
	if idxBiggestPair == -1 {
		fmt.Println("tutup kartu")
		return []int{0, 0}
	}
	fmt.Println(cards[idxBiggestPair])
	return cards[idxBiggestPair]
}

func main() {
	Playingdomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})
}
