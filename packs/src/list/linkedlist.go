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
  cont := 0
  var aux *No = linkedlist.cabeca
  var aux1 *No = aux.prox

  if index == 0{
    linkedlist.cabeca = aux1
  }else{
    for cont < index-1{
      aux = aux.prox
      aux1 = aux.prox
      cont++
    }
    aux.prox = aux1.prox
  }
  linkedlist.tam--
 linkedlist.Print("RemoveOnIndex")
}

func (linkedlist *LinkedList) AddOnIndex(value int, index int){
  cont := 0
  var aux *No = linkedlist.cabeca
  var aux1 *No = aux.prox
  var new_num *No = &No{}

  if index == 0{
    linkedlist.cabeca = new_num
    new_num.value = value
    new_num.prox = aux
  }else{
    for cont < index-1{
      aux = aux.prox
      aux1 = aux.prox
      cont++
    }
    aux.prox = new_num
    new_num.value = value
    new_num.prox = aux1
  }
  linkedlist.tam++
  linkedlist.Print("AddOnIndex")
}

func (linkedlist *LinkedList) Size() int{
  return linkedlist.tam
}

func (linkedlist *LinkedList) Get(index int) int{
  cont := 0
  var aux *No = linkedlist.cabeca
  
  for aux.prox != nil && cont != index{
    aux = aux.prox
    cont++
  }
  return aux.value
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