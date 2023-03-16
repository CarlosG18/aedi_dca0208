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
  linkedlist.cabeca.value = 0
  linkedlist.cabeca.prox = &No{}
  linkedlist.tam = 0
  linkedlist.err = ""
}

func (linkedlist *LinkedList) Add(value int){
  aux := linkedlist.cabeca.prox
  fmt.Println(&aux) 

  if aux.prox == nil{
    aux.value = value
    aux.prox = &No{}
    fmt.Print("colocou o valor")
  }

  //aux1 := aux.prox
  //fmt.Print(&aux1)
/*
  if aux1 != nil{
    fmt.Print("proximo != nil")
  }else{
    fmt.Print("prox == nil")
  }
  
  /*for aux.prox != nil{
    aux = aux.prox
  }

  aux.value = value
  linkedlist.tam++*/
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