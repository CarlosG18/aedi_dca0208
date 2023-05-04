package sort

func QuickSort(v []int, init int, fin int){
  if init < fin{
    pos := Particion(v,init,fin)
    QuickSort(v,init,pos-1)
    QuickSort(v,pos+1,fin)
  }
}

func Particion(v []int, init int, fin int) int{
  pindex := 0
  pivo := v[fin]
  
  for i:=init; i< fin; i++{
    if v[i] <= pivo{
      v[i], v[pindex] = v[pindex], v[i]
      pindex++
    }
  }
  v[pindex], v[fin] = v[fin], v[pindex]
  
  return pindex
}