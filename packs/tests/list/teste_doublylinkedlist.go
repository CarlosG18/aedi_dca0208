package testes

import (
  "fmt"
  "main/packs/src/list"
)

func Teste_DoublyLinkedList(){
  fmt.Println("----- Testando a Doublylinkedlist ----")
  fmt.Println()
  dll := list.DoublyLinkedList{}
  dll.Init()
  for i := 0; i<10; i++{
    dll.Add(i)
  }
  
  dll.Remove()
  
  dll.RemoveOnIndex(3)
  dll.RemoveOnIndex(20)
  dll.RemoveOnIndex(0)
  dll.RemoveOnIndex(-1)
  
  dll.AddOnIndex(150,-1)
  dll.AddOnIndex(150,0)
  dll.AddOnIndex(150,3)
  dll.AddOnIndex(150,30)

  fmt.Println("Size = ", dll.Size())
  fmt.Println("Get = ", dll.Get(-1))
  fmt.Println("Get = ", dll.Get(0))
  fmt.Println("Get = ", dll.Get(3))
  fmt.Println("Get = ", dll.Get(30))

  dll.Set(30,-1)
  dll.Set(30,0)
  dll.Set(30,4)
  dll.Set(30,80)

  for i := 0; i<5; i++{
    dll.AddOnIndex(i,i)
  }
}