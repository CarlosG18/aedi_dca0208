package main

import (
	"fmt"
	"main/packs/src/sort"
)

func main() {
	v := [5]int{45,29,58,36,1}
	fmt.Println(v)
	v1 := sort.InsertionSort(v)
	fmt.Print(v1)

}
