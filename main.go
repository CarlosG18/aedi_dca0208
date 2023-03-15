package main

import (
  "main/packs/list"
)

func main() {
  arraylist := list.ArrayList{}
  arraylist.Init(10)

  for i := 0; i < 10; i++ {
    arraylist.Add(i)
  }

  arraylist.Print()

}
