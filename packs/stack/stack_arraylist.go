package stack

import "errors"

type ArrayStack struct{
  values []int
  size int
}

func (arraystack *ArrayStack) Init(size int){
  arraystack.values = make([]int,size)
  arraystack.size = 0
}

func (arraystack *ArrayStack) Double() {
	doublestack := make([]int, len(arraystack.values)*2)
	for i := 0; i < arraystack.size; i++ {
		doublestack[i] = arraystack.values[i]
	}
	arraystack.values = doublestack
}

func (arraystack *ArrayStack) Push(value int){
	if arraystack.size == len(arraystack.values) {
		arraystack.Double()
	}
	arraystack.values[arraystack.size] = value
	arraystack.size++
}

func (arraystack *ArrayStack) Pop()(int, error){
  if arraystack.size > 0{
    ex_top := arraystack.values[arraystack.size-1]
  arraystack.size--
    return ex_top, nil
  }else{
    return -1, errors.New("a stack esta vazia!")
  }
}

func (arraystack *ArrayStack) Peek() (int,error){
  if arraystack.size > 0{
    return arraystack.values[arraystack.size-1], nil
  }else{
    return -1, errors.New("A stack não possui elementos")
  }
}

func (arraystack *ArrayStack) Size() int{
  return arraystack.size
}

func (arraystack *ArrayStack) Empty() bool{
  if arraystack.size > 0{
    return false
  }else{
    return true
  }
}