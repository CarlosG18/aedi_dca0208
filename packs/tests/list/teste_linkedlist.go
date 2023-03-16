package testes

import "main/packs/src/list"

func Teste_linkedlist(){
  linkedlist := list.LinkedList{}
  linkedlist.Init()
  linkedlist.Add(67)
  linkedlist.Add(89)
}