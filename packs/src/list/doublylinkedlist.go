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
  var aux *No1 = doublylinkedlist.head

  for aux.prox != nil{
    aux = aux.prox
  }
  aux.value = value
  aux.prox = &No1{}
  aux.prox.ant = aux

  doublylinkedlist.tam++
  doublylinkedlist.Print("Add")
}

func (doublylinkedlist *DoublyLinkedList) Remove(){
  var aux *No1 = doublylinkedlist.head
  for aux.prox != nil{
    aux = aux.prox
  }
  aux.ant.prox = nil
  
  doublylinkedlist.tam--
  doublylinkedlist.Print("Remove")
}

func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int){
  var cont = 0
  var aux *No1 = doublylinkedlist.head
  
  for cont != index{
    aux = aux.prox
    cont++
  }
  aux.ant.prox = aux.prox

  doublylinkedlist.tam--
  doublylinkedlist.Print("RemoveOnIndex")
}

func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value int, index int){
  var cont = 0
  var aux *No1 = doublylinkedlist.head
  
  for cont != index{
    aux = aux.prox
    cont++
  }
  aux1 := &No1{}
  aux.ant.prox = aux1
  aux1.value = value
  aux1.ant = aux.ant
  aux1.prox = aux

  doublylinkedlist.tam++
  doublylinkedlist.Print("AddOnIndex")
}

func (doublylinkedlist *DoublyLinkedList) Size() int{
  return 0
}

func (doublylinkedlist *DoublyLinkedList) Get(index int) int{
  return 0
}
 
func (doublylinkedlist *DoublyLinkedList) Set(value int, index int){

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
}