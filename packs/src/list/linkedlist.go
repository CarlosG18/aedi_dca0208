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
  linkedlist.Reset_err()
  
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
  linkedlist.Reset_err()
  var aux *No = linkedlist.cabeca
  var auxAnt *No

  if linkedlist.tam > 0{
    for aux.prox != nil {
    auxAnt = aux
    aux = aux.prox
    }
    if aux.prox == nil{
    auxAnt.prox = nil
    }
    linkedlist.tam--
    linkedlist.Print("Remove")
  }else{
    linkedlist.err = "error:A linkedlist não possui elementos!"
  linkedlist.Print("Remove:error")
  }
}

func (linkedlist *LinkedList) RemoveOnIndex(index int){
  linkedlist.Reset_err()
  cont := 0
  var aux *No = linkedlist.cabeca
  var aux1 *No = aux.prox

  if linkedlist.tam > 0 && (index < linkedlist.tam && index >= 0) {
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
  }else{
    linkedlist.err = "error: Não foi possível acessar o index!"
    linkedlist.Print("RemoveOnIndex:error")
  }
}

func (linkedlist *LinkedList) AddOnIndex(value int, index int){
  linkedlist.Reset_err()
  cont := 0
  var aux *No = linkedlist.cabeca
  var aux1 *No = aux.prox
  var new_num *No = &No{}

  if index < 0 || index >= linkedlist.tam {
    linkedlist.err = "error: não foi possivel acessar o index!"
    linkedlist.Print("AddOnIndex:error")
  }else{
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
}

func (linkedlist *LinkedList) Size() int{
  linkedlist.Reset_err()
  return linkedlist.tam
}

func (linkedlist *LinkedList) Get(index int) int{
  linkedlist.Reset_err()
  cont := 0
  var aux *No = linkedlist.cabeca

  if index < 0 || index >= linkedlist.tam {
    linkedlist.err = "error: não foi possivel acessar o index!"
    linkedlist.Print("Get:error")
    return -1
  }else{
    for aux.prox != nil && cont != index{
      aux = aux.prox
      cont++
    }
  linkedlist.Print("Get")
  return aux.value
  }
}
 
func (linkedlist *LinkedList) Set(value int, index int){
  linkedlist.Reset_err()
  cont := 0
  var aux *No = linkedlist.cabeca

  if index < 0 || index >= linkedlist.tam {
    linkedlist.err = "error: não foi possivel acessar o index!"
    linkedlist.Print("Set:error")
  }else{
    for aux.prox != nil && cont != index{
      aux = aux.prox
      cont++
    }
  aux.value = value
  linkedlist.Print("Set")
  }
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
  if linkedlist.err != ""{
    fmt.Println(linkedlist.err)
  }
  fmt.Println("tamanho da linkedlist = ", linkedlist.tam)
  fmt.Println()
}

func (linkedlist *LinkedList) Reset_err(){
  linkedlist.err = ""
}