package main

import (
	"fmt"
	"main/packs/src/queue"
)

func main() {
	fila := queue.QueueArray{}
	fila.Init(10)
  
	for i := 0; i < 10; i++ {
		fila.Enqueue(i)
	}
	fmt.Println()
	fmt.Println("desenfileirando")
	fmt.Println()
	for i := 0; i < 10; i++ {
		fila.Dequeue()
	}

  for i := 0; i < 10; i++ {
		fila.Enqueue(i)
	}
  
  
}
