package Queue

type QueueArray struct{
  values []int
  size int
  front int
  near int
}

func (queue *QueueArray) Enqueue(value int){
  
}

func (queue *QueueArray) Dequeue() (int, error){
  return 0, nil
}

func (queue *QueueArray) Size() int{
  return 0
}


func (queue *QueueArray) Front() (int, error){
  return 0, nil
}
  