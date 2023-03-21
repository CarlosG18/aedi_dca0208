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
  dll.Add(56)
  dll.Add(90)
 /* for i := 0; i<10; i++{
    dll.Add(i)
  }*/
  
  
}