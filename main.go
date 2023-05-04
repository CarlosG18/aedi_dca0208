package main

import (
	"fmt"
	"main/packs/src/sort"
	"math/rand"
	//"time"
)

func main() {
  vetor := FullArrayRandom(5)
  fmt.Println("vetor = ", vetor)
  //start := time.Now()
  vetor_ord := sort.MergeSort(vetor, len(vetor))
  //end := time.Since(start)
  fmt.Println("vetor ordenado com Mergesort = ", vetor_ord)
  //fmt.Println()
  //fmt.Println("tempo gasto = ", end)
}

func FullArrayRandom(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = rand.Intn(100)
  }
  return vetor
}