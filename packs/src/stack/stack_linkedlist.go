package stack

//import "errors"

type Stack struct{
  head *No
  size int
}

type No struct{
  value int
  prox *No
}

func (stack *Stack) Push(value int){
  
}

func (stack *Stack) Pop(){
  
}

func (stack *Stack) Top() int{
  return -1
}

func (stack *Stack) Size() int{
  return -1
}

func (stack *Stack) Empty() bool{
  return true
}