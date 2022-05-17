package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxminBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
	fmt.Println(maxminBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}))
}

func maxminBuyProduct(money int, product []int) int {
	sort.Ints(product)
	var i = 0
	var countProduct = 0
	for money >= 0 && i < len(product) {
		money -= product[i]
		countProduct++
		i++
	}
	return countProduct - 1
}
