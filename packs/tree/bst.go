package tree

import "fmt"

type Bst struct{
  root *BstNode
  size int
}

type BstNode struct{
  left *BstNode
  value int
  rigth *BstNode
}

func Print(nome string){
  fmt.Println(nome)  
}

func (bst *Bst) Init(){
  bst.root = &(BstNode{})
  bst.size = 0
}

func NewNode(value int) *BstNode{
  node := BstNode{}
  node.left = nil
  node.rigth = nil
  node.value = value
  return &node
}

func (bst *Bst) Add(value int){
  if bst.root == nil{
    bst.root = NewNode(value)
  }
  aux := bst.root
  prev := aux
  for aux != nil{
    prev = aux
    if value <= aux.value{
      aux = aux.left 
    }else{
      aux = aux.rigth
    }    
  }
  if value <= prev.value{
    prev.left = NewNode(value)
  }else{
    prev.rigth = NewNode(value)
  }
}

func (bst *Bst) Busca(value int) bool{
  aux := bst.root

  for aux != nil{
    if value == aux.value{
      return true
    }else{
      if value < aux.value{
        aux = aux.left
      }else{
        aux = aux.rigth
      }
    }
  }
  return false
}

// func (bst *Bst) Print(){
//   aux := bst.root
//   for aux != nil{
//    fmt.Println("Raiz = ", aux.value)
//     fmt.Println("valores a esquerda")
//     for aux.left != nil{
//       aux = aux.left
//       fmt.Println("valor = ", aux.value)
//     }
//     fmt.Println("valores a direita")
//     if aux.rigth != nil{
//       aux = aux.rigth
//       fmt.Println("valor = ", aux.value)
//     }
//   }
// }

