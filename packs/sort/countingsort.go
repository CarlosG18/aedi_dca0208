package sort

func CountingSort(v []int) []int{
  /*
  passos:
  1 descobrir o maior e menor elemento do vetor
  2 criar um vetor de contagem com tam = maior - menor + 1
  3 contar as ocorrências de cada elemento em v
  4 somar comulativamente no vetor de contagem
  5 alocar um novo vetor para armazenar os valores ordenados
  6 mapear os elementos de v para o usando c
  */
  
  tam := len(v)
  maior := v[0]
  menor := v[0]
  
  for i:=1; i<tam; i++{
    if v[i] > maior{
      maior = v[i]
    }
    if v[i] < menor{
      menor = v[i]
    }
  }
  
  c := make([]int,maior-menor+1)
  
  for i:=0; i < len(v); i++{
    c[v[i]-menor]++
  }
  //soma acumulativa em c
  for i:= 1; i < len(c); i++{
    c[i] += c[i-1]
  }
  
  o := make([]int,tam)
  for i:=0; i< len(v); i++{
    o[c[v[i]-menor]-1] = v[i]
  }
  
  return o
}