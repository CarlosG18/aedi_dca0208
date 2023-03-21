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

}

func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int){
  
}

func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value int, index int){
  
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
  fmt.Println(doublylinkedlist.err)
}