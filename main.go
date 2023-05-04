package main

import (
	//"fmt"
	"main/packs/src/sort"
)

func main() {
  vetor := [7]int{2,7,9,12,0,1,5}
  sort.QuickSort(vetor,0,len(vetor)-1)
}
