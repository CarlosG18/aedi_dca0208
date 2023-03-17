package list

import "fmt"

type LinkedList struct{
  cabeca *No
  tam int
  err string
}

type No struct{
  value int
  prox *No
}

func (linkedlist *LinkedList) Init(){
  linkedlist.cabeca = &No{}
  linkedlist.tam = 0
  linkedlist.err = ""
}

func (linkedlist *LinkedList) Add(value int){
  var aux *No = linkedlist.cabeca
  
  for aux.prox != nil {
    aux = aux.prox
  }
  if aux.prox == nil{
    aux.value = value
    aux.prox = &No{}
  }
  linkedlist.tam++
  linkedlist.Print("Add")
}

func (linkedlist *LinkedList) Remove(){
  var aux *No = linkedlist.cabeca
  var auxAnt *No
  
  for aux.prox != nil {
    auxAnt = aux
    aux = aux.prox
  }
  if aux.prox == nil{
    auxAnt.prox = nil
  }
  linkedlist.tam--
  linkedlist.Print("Remove")
}

func (linkedlist *LinkedList) RemoveOnIndex(index int){
  
}

func (linkedlist *LinkedList) AddOnIndex(value int, index int){
  
}

func (linkedlist *LinkedList) Size() int{
  return linkedlist.tam
}

func (linkedlist *LinkedList) Get(index int) int{
  return 1
}
 
func (linkedlist *LinkedList) Set(value int, index int){
  
}

func (linkedlist *LinkedList) Print(operation string){
  var aux *No = linkedlist.cabeca
  fmt.Print("[")
  for aux.prox != nil {
    fmt.Print(aux.value)
    aux = aux.prox
    if aux.prox != nil{
      fmt.Print(",")
    }
  }
  fmt.Println("]")
  fmt.Println(operation)
  fmt.Println("tamanho da linkedlist = ", linkedlist.tam)
  fmt.Println()
}