package queue

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
  Empty() bool
  Full() bool
}
