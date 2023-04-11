package queue

import (
	"errors"
	"fmt"
)

type QueueArray struct {
	values []int
	size   int
	front  int
	near   int
}

func (queue *QueueArray) Init(size int) {
	queue.front = -1
	queue.near = -1
	queue.values = make([]int, size)
}

func (queue *QueueArray) Enqueue(value int) {
  queue.front = 0
  queue.near = 
  queue.values[] = value
	queue.size++
}

func (queue *QueueArray) Dequeue() (int, error) {
  /*
	if queue.front < queue.near {
		queue.front++
		queue.size--
		queue.Print()
		return queue.values[queue.front], nil
	} else {
		queue.front = -1
		queue.near = -1
		queue.Print()
		return -1, errors.New("a fila nao possui elementos")
	}*/
}

func (queue *QueueArray) Size() int {
	return queue.size
}

func (queue *QueueArray) Front() (int, error) {
	if queue.front < queue.near {
		return queue.front, nil
	} else {
		return -1, errors.New("a fila esta vazia!")
	}
}

func (queue *QueueArray) Print() {
	if queue.front != -1 && queue.near != -1 {
		for i := queue.front; i <= queue.near; i++ {
			fmt.Print(queue.values[i], " ")
		}
	}
	fmt.Println("front - ", queue.front, " end - ", queue.near)
}
