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
  dll.RemoveOnIndex(4)
  
}