package main

import "fmt"

func fibonanci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonanci(n-1) + fibonanci(n-2)
}
func main() {
	fmt.Println(fibonanci(0))
	fmt.Println(fibonanci(1))
	fmt.Println(fibonanci(2))
	fmt.Println(fibonanci(3))
	fmt.Println(fibonanci(4))
	fmt.Println(fibonanci(5))
}
