package sort

func MergeSort(v []int){
  //função para dividir o vetor v recursivamente
}

func Merge(v1 []int, v2 []int) []int{
  n := len(v1)
  m := len(v2)
  new_v := make([]int,n+m)
  cont1 := 0
  cont2 := 0
  cont3 := 0
  
  for cont3 < len(new_v){
    if(v1[cont1] > v2[cont2]){
      new_v[i] = v2[cont2]
      cont3++
      cont2++
    }else{
      new_v[i] = v1[cont1]
      cont3++
      cont1++
    }
  }
  
  return new_v
}