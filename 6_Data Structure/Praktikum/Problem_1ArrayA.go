package main

import "fmt"

func main() {
	///arrayA
	arraySatu := []string{"kazuya", "jin", "lee"}
	arrayDua := []string{"kazuya", "feng"}
	fmt.Println(arrayMerge(arraySatu, arrayDua))
}
func arrayMerge(arraySatu, arrayDua []string) (hasilGabung []string) {
	hmap := make(map[string]bool)
	for _, ValueSatu := range arraySatu {
		hmap[ValueSatu] = true
		hasilGabung = append(hasilGabung, ValueSatu)
	}
	for _, ValueDua := range arrayDua {
		_, isFound := hmap[ValueDua]
		if isFound == false {
			hmap[ValueDua] = true
			hasilGabung = append(hasilGabung, ValueDua)
		}
	}
	return hasilGabung
}
