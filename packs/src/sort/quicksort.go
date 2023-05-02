package sort

func QuickSort(v []int, init int, fin int){
  if init < fin{
    pos = Particion(v,init,fin
    QuickSort(v,init,pos-1)
    QuickSort(v,pos+1,fin)
  }
}

func Particion(v []int, init int, fin int) int{
  tam := len(v)
  i := 0
  pindex := 0
  pivo := v[fin]
  
  for i < tam{
    if v[i] < pivo{
      i++
    }else{
      pindex++
      
    }
  }
}