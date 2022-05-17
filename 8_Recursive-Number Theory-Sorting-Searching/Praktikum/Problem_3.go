package main

import "fmt"

func PrimaSegiEmpat(wide, height, start int) int {
	total := 0
	for i := 0; i < height; i++ {
		j := 0
		for j < wide {
			start++
			if val, ok := isPrime(start); ok {
				fmt.Printf("%v\t", val)
				total += val
				j++
			}
		}
		fmt.Println("\n")
	}
	return total
}

func isPrime(start int) (int, bool) {
	if start <= 1 {
		return 0, false
	}
	for i := 2; i < start; i++ {
		if start%i == 0 {
			return 0, false
		}
	}
	return start, true
}

func main() {
	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
}
