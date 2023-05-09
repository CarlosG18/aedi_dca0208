package main

import (
	"fmt"
	"main/packs/sort"
	"math/rand"
	"time"
)

func main() {
  
  vetor := FullArrayRandom(30)
  fmt.Println("vetor = ", vetor)
  fmt.Println()
  start := time.Now()
  vetor_ord := sort.CountingSort(vetor)
  end := time.Since(start)
  fmt.Println("vetor ordenado com CountingSort= ", vetor_ord)
  fmt.Println("tempo gasto com CountingSort = ", end)
  
  fmt.Println()
  /*
  start1 := time.Now()
  sort.QuickSort(vetor,0,len(vetor)-1)
  end1 := time.Since(start1)
  fmt.Println("vetor ordenado com QuickSort = ", vetor)
  fmt.Println("tempo gasto com QuickSort = ", end1)*/
}

func FullArrayRandom(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = rand.Intn(100)
  }
  return vetor
}

func FullArraySalary(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = i
  }
  return vetor
}