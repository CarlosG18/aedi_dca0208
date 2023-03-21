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
  doublylinkedlist.head.prox = &No1{}
  doublylinkedlist.tam = 0
  doublylinkedlist.err = ""
}

func (doublylinkedlist *DoublyLinkedList) Add(value int){

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
  
  fmt.Println("Tamanho da doublylinkedlist = ", doublylinkedlist.tam)
  fmt.Println(doublylinkedlist.err)
}