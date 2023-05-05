package stack

import "errors"

type LinkedStack struct{
  head *No
  size int
}

type No struct{
  value int
  prox *No
}

func (linkedstack *LinkedStack) Push(value int){
  new_num := &No{}
  aux := linkedstack.head

  linkedstack.head = new_num
  new_num.value = value
  new_num.prox = aux
  linkedstack.size++
}

func (linkedstack *LinkedStack) Pop() (int,error){
  aux := linkedstack.head
  if linkedstack.size > 0{
    linkedstack.head = aux.prox
    linkedstack.size--
    return aux.value, nil
  }else{
    return -1, errors.New("a stack não possui elementos!")
  }
}

func (linkedstack *LinkedStack) Peek() (int,error){
  if linkedstack.size > 0{
    return linkedstack.head.value, nil  
  }else{
    return 1, errors.New("a stack não possui elementos!")  
  }
}

func (linkedstack *LinkedStack) Size() int{
  return linkedstack.size
}

func (linkedstack *LinkedStack) Empty() bool{
  if linkedstack.size > 0{
    return false
  }else{
    return true
  }
}