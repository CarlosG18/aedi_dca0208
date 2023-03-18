package testes

import (
  "fmt"
  "main/packs/src/list")

func Teste_arraylist(){
  fmt.Println(" ------ TESTANDO ARRAYLIST ------")
  
  arraylist := list.ArrayList{}
  arraylist.Init(10)
  for i:=0; i<10; i++{
    arraylist.Add(i)
  }
  arraylist.RemoveOnIndex(200)
  arraylist.Remove()
  arraylist.AddOnIndex(59,2)
  arraylist.Set(1,100)
  arraylist.AddOnIndex(89,100)
}