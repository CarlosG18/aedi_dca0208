package main

import (
	"fmt"
	"main/packs/src/queue"
)

func main() {
	fila := queue.QueueArray{}
	fila.Init()
  
	for i := 0; i < 5; i++ {
		fila.Enqueue(i)
  }
  
	fmt.Println()
	fmt.Println("desenfileirando")
	fmt.Println()
  
	for i := 0; i < 2; i++ {
		fila.Dequeue()
	}
  
  for i := 0; i < 8; i++ {
		fila.Enqueue(i)
	}

  for i := 0; i < 11; i++ {
		fila.Dequeue()
	}

  
}
