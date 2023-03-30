package stack

import "errors"

type Stack struct{
  values []int
  size int
}

func (stack *Stack) Init(size int){
  stack.values = make([]int,size)
  stack.size = 0
}

func (stack *Stack) Double() {
	doublestack := make([]int, len(stack.values)*2)
	for i := 0; i < stack.size; i++ {
		stack[i] = stack.values[i]
	}
	stack.values = doublestack
}

func (stack *Stack) Push(value int){
	if stack.size == len(stack.values) {
		stack.Double()
	}
	stack.values[stack.size] = value
	stack.size++
}

func (stack *Stack) Pop(){
  stack.size--
}

func (stack *Stack) Top() int{
	return arraylist.values[stack.size]
}

func (stack *Stack) Size() int{
  return stack.size
}

func (stack *Stack) Empty() bool{
  if stack.size == 0{
    return false
  }else{
    return true
  }
}