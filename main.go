package main

import (
	"fmt"
  "main/packs/src/sort"
)

func main() {
	v := []int{45,298,5,89,1}
	fmt.Println(v)
	sort.MergeSort(v, len(v))
	/*v1 := sort.MergeSort(v, len(v))
	fmt.Print(v1)*/
	
}
