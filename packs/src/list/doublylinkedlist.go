package list

import "fmt"

type DoublyLinkedList struct{
  head *No1
  tam int
  err string
}

type No1 struct{
  value int 
  ant *No1
  prox *No1
}

func (doublylinkedlist *DoublyLinkedList) Init(){
  doublylinkedlist.head = &No1{}
  doublylinkedlist.head.ant = nil
  doublylinkedlist.tam = 0
  doublylinkedlist.err = ""
}

func (doublylinkedlist *DoublyLinkedList) Add(value int){
  doublylinkedlist.err = ""
  aux := doublylinkedlist.Get_no_nil()
  aux.value = value
  aux.prox = &No1{}
  aux.prox.ant = aux

  doublylinkedlist.tam++
  doublylinkedlist.Print("Add")
}

func (doublylinkedlist *DoublyLinkedList) Remove(){
  doublylinkedlist.err = ""
  aux := doublylinkedlist.Get_no_nil()
  aux.ant.prox = nil
  
  doublylinkedlist.tam--
  doublylinkedlist.Print("Remove")
}

func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int){
  doublylinkedlist.err = ""
  if index == 0{
    aux := doublylinkedlist.head
    doublylinkedlist.head = aux.prox
    aux.prox.ant = nil
  }else{
    if index >=0 && index < doublylinkedlist.tam{
      aux := doublylinkedlist.Get_no_cont(index)
      aux.ant.prox = aux.prox
      doublylinkedlist.tam--
    }else{
      doublylinkedlist.err = "error: não foi possivel acessar o index!"
    }
  }
  doublylinkedlist.Print("RemoveOnIndex")
}

func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value int, index int){
  doublylinkedlist.err = ""
  if index == 0{
    aux1 := &No1{}
    aux := doublylinkedlist.head
    doublylinkedlist.head = aux1
    aux1.prox = aux
    aux.ant = aux1
    aux1.value = value
    aux1.ant = nil
    doublylinkedlist.tam++
  }else if index > 0 && index < doublylinkedlist.tam{
    aux := doublylinkedlist.Get_no_cont(index)
  fmt.Print("index = ", index)
  aux1 := &No1{}
  aux.ant.prox = aux1
  aux1.value = value
  aux1.ant = aux.ant
  aux1.prox = aux
  aux.ant = aux1
  doublylinkedlist.tam++
  }else{
    doublylinkedlist.err = "error: não foi possivel acessar o index!"
  }
  doublylinkedlist.Print("AddOnIndex")
}

func (doublylinkedlist *DoublyLinkedList) Size() int{
  doublylinkedlist.err = ""
  doublylinkedlist.Print("Size")
  return doublylinkedlist.tam
}

func (doublylinkedlist *DoublyLinkedList) Get(index int) int{
  aux := doublylinkedlist.head
  doublylinkedlist.err = ""
  if index >= 0 && index < doublylinkedlist.tam{
    aux = doublylinkedlist.Get_no_cont(index)
  }else{
    doublylinkedlist.err = "error: não foi possivel acessar o index!"
    doublylinkedlist.Print("Get")
    return -1
  }
  
  doublylinkedlist.Print("Get")
  return aux.value
}
 
func (doublylinkedlist *DoublyLinkedList) Set(value int, index int){
  doublylinkedlist.err = ""
  if index >= 0 && index < doublylinkedlist.tam{
    aux := doublylinkedlist.Get_no_cont(index)
    aux.value = value
  }else{
    doublylinkedlist.err = "error: não foi possivel acessar o index!"
  }
  
  doublylinkedlist.Print("Set")
}

func (doublylinkedlist *DoublyLinkedList) Print(operation string){
  var aux *No1 = doublylinkedlist.head
  fmt.Print("[")
  for aux.prox != nil{
    fmt.Print(aux.value)
    aux = aux.prox
    if aux.prox != nil{
      fmt.Print(",")
    }
  }
  fmt.Print("]")
  fmt.Println()
  fmt.Println("Tamanho da doublylinkedlist = ", doublylinkedlist.tam)
  fmt.Println(operation)
  fmt.Println(doublylinkedlist.err)
  fmt.Println()
}

func (doublylinkedlist *DoublyLinkedList) Get_no_nil() *No1{
  var aux *No1 = doublylinkedlist.head
  for aux.prox != nil{
    aux = aux.prox
  }
  return aux
}

func (doublylinkedlist *DoublyLinkedList) Get_no_cont(index int) *No1{
  var cont = 0
  var aux *No1 = doublylinkedlist.head
  for cont != index{
    aux = aux.prox
    cont++
  }
  return aux
}
