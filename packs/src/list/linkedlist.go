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
  
  
  linkedlist.cabeca.value = value
  linkedlist.cabeca.prox = &No{}
  linkedlist.tam++
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
  fmt.Print("[")
  for i := 0; i < linkedlist.tam; i++{
    
  }
}