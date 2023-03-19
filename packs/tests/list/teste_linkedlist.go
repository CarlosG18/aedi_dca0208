package testes

import (
  "fmt"
  "main/packs/src/list"
)

func Teste_linkedlist(){
  fmt.Println("----- Testando a linkedlist ----")
  fmt.Println()

  linkedlist := list.LinkedList{}
  linkedlist.Init()
  
  for i := 0; i<10; i++{
    linkedlist.Add(i)
  }

  linkedlist.RemoveOnIndex(-1)
  linkedlist.RemoveOnIndex(10)
  linkedlist.AddOnIndex(500,-1)
  linkedlist.AddOnIndex(78,44)
  fmt.Println(linkedlist.Get(-1))
  fmt.Println(linkedlist.Get(22))
  linkedlist.Set(45,56)
  linkedlist.Set(45,-1)
}