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
    aux.prox = &No{}
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
  
}

func (linkedlist *LinkedList) RemoveOnIndex(index int){
  
}

func (linkedlist *LinkedList) AddOnIndex(value int, index int){
  
}

func (linkedlist *LinkedList) Size(){
  
}

func (linkedlist *LinkedList) Get(index int){
  
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
  fmt.Println()
}