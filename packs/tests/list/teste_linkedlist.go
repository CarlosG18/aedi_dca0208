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

  for i := 0; i<5; i++{
    linkedlist.Remove()
  }

  linkedlist.AddOnIndex(200,3)
  linkedlist.AddOnIndex(150,1)
  linkedlist.AddOnIndex(55,0)
  linkedlist.AddOnIndex(120,0)
  linkedlist.RemoveOnIndex(0)
  linkedlist.RemoveOnIndex(2)
  fmt.Println(linkedlist.Get(0))
}