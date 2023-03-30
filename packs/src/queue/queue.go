package Queue

type IQueue interface {
	Enqueue(value int) 
  Dequeue() (int, error)
  Size() int 
  Front() int
}