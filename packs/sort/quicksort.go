package sort

//import "math/rand"

func QuickSort(v []int, init int, fin int){
  if init < fin{
    pos := Particion(v,init,fin)
    QuickSort(v,init,pos-1)
    QuickSort(v,pos+1,fin)
  }
}

func Particion(v []int, init int, fin int) int{
  pindex := init
  /*i := rand.Intn(fin)
  v[fin], v[i] = v[i], v[fin]*/
  pivo := v[fin]
  
  for i:=init; i< fin; i++{
    if pivo >= v[i] {
      v[i], v[pindex] = v[pindex], v[i]
      pindex++
    }
  }
  v[pindex], v[fin] = v[fin], v[pindex]
  
  return pindex
}