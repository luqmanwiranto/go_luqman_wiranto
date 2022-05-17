package main

import "fmt"

func main() {
	fmt.Println(Eksponen(2, 3))
	fmt.Println(Eksponen(7, 2))
}

func Eksponen(bilangan, pangkat int) (hasilEksponen int) {
	if pangkat == 0 {
		hasilEksponen = 1
		return hasilEksponen
	}
	hasilEksponen = bilangan
	for i := 1; 1 <= pangkat; i++ {
		hasilEksponen = hasilEksponen * bilangan
	}
	return hasilEksponen
}
