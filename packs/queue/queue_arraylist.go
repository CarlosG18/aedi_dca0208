package queue

import (
	//"errors"
	"fmt"
)

type QueueArray struct {
	values []int
	size   int
  n int
	front  int
	near   int
}

func (queue *QueueArray) Init() {
	queue.front = -1
	queue.near = -1
  queue.n = 10
  queue.values = make([]int,queue.n)
}

func (queue *QueueArray) Enqueue(value int) {
  if queue.front == -1{
    queue.front = 0
  }

  if ((queue.near + 1) % queue.n) == queue.front && (queue.front != 0 && queue.near != 0){
    fmt.Println("fila cheia - segundo o prof")
  }else{
    queue.near = (queue.near+1)%queue.n
    queue.values[queue.near] = value
    queue.size++
    queue.Print()
  }
}

func (queue *QueueArray) Dequeue() int {
  if queue.Empty(){
    fmt.Println("fila vazia!")
    queue.front = -1
    queue.near = -1
    return -1
  }else{
    retorno := queue.values[queue.front]
    queue.front = (queue.front+1)%queue.n
    queue.size--
    queue.Print()
    return retorno
  }
}

func (queue *QueueArray) Empty() bool {
	return queue.size==0
}

func (queue *QueueArray) Print() {
	fmt.Println("front - ", queue.front, " end - ", queue.near)
}
