package main

import (
	"fmt"
	"main/packs/src/sort"
)

func main() {
	v := []int{298, 45, 5, 89, 1}
	fmt.Println(v)
	v1 := sort.MergeSort(v, len(v))
	fmt.Print(v1)
}
