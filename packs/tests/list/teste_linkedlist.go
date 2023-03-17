package testes

import "main/packs/src/list"

func Teste_linkedlist(){
  linkedlist := list.LinkedList{}
  linkedlist.Init()
  for i := 0; i<10; i++{
    linkedlist.Add(i)
  }
}