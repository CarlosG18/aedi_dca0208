package main

import (
	//"fmt"
	"main/packs/src/list"
)

func main() {
	link := list.LinkedList{}
	link.Init()

	for i := 0; i < 10; i++ {
		link.Add(i)
	}
	link.Print()
	link.Invert()
	link.Print()
  
}
