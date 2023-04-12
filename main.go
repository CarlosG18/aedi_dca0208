package main

import (
	"fmt"
	"main/packs/src/list"
)

func main() {
	ll := list.LinkedList{}
	ll.Init()

	ll.Add(51)
  ll.Add(1)
  ll.Add(4)
  ll.Add(6)
  ll.Add(5)

	for i := 0; i < 10; i++ {
		ll.Add(i)
	}

	ll.Print()
  
	fmt.Println()
	ll.SubList_Par()
	ll.Print()
  
}
