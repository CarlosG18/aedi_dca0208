package queue

import "errors"

type QueueArray struct{
  values []int
  size int
  front int
  near int
}

func (queue *QueueArray) Init(size int){
  queue.front = -1
  queue.near = -1
  queue.values = make([]int,size)
}

func (queue *QueueArray) Double(){
  double_queue := make([]int,queue.size*2)
  for i := 0; i<queue.size-1; i++{
    double_queue[i] = queue.values[i]
  }
  queue.values = double_queue
}

func (queue *QueueArray) Enqueue(value int){
  if queue.size == len(queue.values){
    queue.Double()
  }
  queue.values[queue.size] = value
  queue.front = 0
  queue.near = queue.size
  queue.size++
}

func (queue *QueueArray) Dequeue() (int, error){
  if queue.front < queue.near{
    queue.front++
    queue.size--
  }else{
    queue.front = -1
    queue.near = -1
    return -1, errors.New("a fila nao possui elementos")
  }
}

func (queue *QueueArray) Size() int{
  return queue.size
}


func (queue *QueueArray) Front() (int, error){
  if queue.front < queue.near{
    return queue.front, nil
  }else{
    return -1, errors.New("a fila esta vazia!")
  }
}
  