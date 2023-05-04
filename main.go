package main

import (
	"fmt"
	"main/packs/src/sort"
	"math/rand"
)

func main() {
  vetor := FullArrayRandom(23)
  fmt.Println("vetor = ", vetor)
  sort.QuickSort(vetor,0,len(vetor)-1)
  fmt.Println("vetor ordenado com quicksort = ", vetor)
}

func FullArrayRandom(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = rand.Intn(100)
  }
  return vetor
}