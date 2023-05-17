package main

import (
	"fmt"
	//"main/packs/sort"
  "main/packs/tree"
	"math/rand"
	//"time"
)

func main() {
  /* testes algortitmos de ordenação
  vetor := FullArrayDecres(10000)
  fmt.Println("vetor = ", vetor)
  fmt.Println()
  
  //BubbleSort
  start := time.Now()
  vetor_bubble := sort.BubbleSort(vetor)
  end := time.Since(start)
  fmt.Println("vetor ordenado com BubbleSort = ", vetor_bubble)
  fmt.Println()
  
  //SelectionSort
  start1 := time.Now()
  vetor_selection := sort.SelectionSort(vetor)
  end1 := time.Since(start1)
  fmt.Println("vetor ordenado com SelectionSort = ", vetor_selection)
  
  //InsertionSort
  start2 := time.Now()
  vetor_insertion := sort.InsertionSort(vetor)
  end2 := time.Since(start2)
  fmt.Println("vetor ordenado com InsertionSort = ", vetor_insertion)
  
  //MergeSort
  start3 := time.Now()
  vetor_merge := sort.MergeSort(vetor,len(vetor))
  end3 := time.Since(start3)
  fmt.Println("vetor ordenado com MergeSort = ", vetor_merge)
  
  vetor_quick := vetor
  
  //QuickSort
  start4 := time.Now()
  sort.QuickSort(vetor_quick,0,len(vetor_quick)-1)
  end4 := time.Since(start4)
  fmt.Println("vetor ordenado com QuickSort = ", vetor_quick)
  
  //CountingSort
  start5 := time.Now()
  vetor_couting := sort.CountingSort(vetor)
  end5 := time.Since(start5)
  fmt.Println("vetor ordenado com CountingSort = ", vetor_couting)
  fmt.Println()
  
  fmt.Println("tempo gasto com BubbleSort O(n²) = ", end)
  fmt.Println("tempo gasto com SelectionSort O(n²) = ", end1)
  fmt.Println("tempo gasto com InsertionSort O(n²) = ", end2)
  fmt.Println("tempo gasto com MergeSort O(logn) = ", end3)
  fmt.Println("tempo gasto com QuickSort O(nlogn) = ", end4)
  fmt.Println("tempo gasto com CountingSort O(n) = ", end5)
*/

// fazendo testes em arvores
  for i:=0; i<10; i++{
    TesteBusca(i)
  }
}

func TesteBusca(numero int){
  tree := tree.Bst{}
  vetor_numeros := FullArrayRandom(10,10)
  fmt.Println("imprimindo o vetor de valores = ", vetor_numeros)

  for i:=0; i<len(vetor_numeros); i++{
    tree.Add(vetor_numeros[i])
  }

  fmt.Println("procurando o numero ", numero)
  buscar := tree.Busca(numero)
  fmt.Println("achou? ", buscar)
  fmt.Println()
}

func FullArrayRandom(tam int, intervalo int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = rand.Intn(intervalo)
  }
  return vetor
}

func FullArrayCres(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = i
  }
  return vetor
}

func FullArrayDecres(tam int) []int{
  vetor := make([]int,tam)
  
  for i:=0; i<tam; i++{
    vetor[i] = (tam-1)-i
  }
  return vetor
}