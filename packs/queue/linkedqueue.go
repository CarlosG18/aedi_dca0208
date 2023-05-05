package queue

type LinkedQueue struct{
  head *No
  size int
  front int
  near int
}

type No struct{
  value int
  prox *No
}

func (linkedqueue *LinkedQueue) Enqueue(value int){
  
}

func (linkedqueue *LinkedQueue) Dequeue() (int, error){
  return 0, nil
}

func (linkedqueue *LinkedQueue) Size() int{
  return 0
}


func (linkedqueue *LinkedQueue) Front() (int, error){
  return 0, nil
}