package tree

import "fmt"

type BstNode struct{
  left *BstNode
  value int
  rigth *BstNode
}

func NewNode(value int) *BstNode{
  node := BstNode{}
  node.left = nil
  node.rigth = nil
  node.value = value
  return &node
}

func (bstnode *BstNode) Add(value int){
  if value > bstnode.value{
    if bstnode.rigth != nil{
      bstnode.rigth.Add(value)
    }else{
      bstnode.rigth = NewNode(value)
    }
  }else{
    if bstnode.left != nil{
      bstnode.left.Add(value)
    }else{
      bstnode.left = NewNode(value)
    }
  }
}

func (bstnode *BstNode) Search(value int) bool{
  if value == bstnode.value{
    return true
  }else if value > bstnode.value{
    if bstnode.rigth == nil{
      return false
    }else{
      return bstnode.rigth.Search(value)
    }
  }else{
    if bstnode.left == nil{
      return false
    }else{
      return bstnode.left.Search(value)
    }
  }
}

//navegação em arvores

func (bstnode *BstNode) PrintIn(){
  if bstnode.left != nil{
    bstnode.left.PrintIn()
  }
  fmt.Print(bstnode.value, ", ")
  if bstnode.rigth != nil{
    bstnode.rigth.PrintIn()
  }
}

func (bstnode *BstNode) PrintPre(){
  fmt.Print(", ", bstnode.value)
  if bstnode.left != nil{
    bstnode.left.PrintIn()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintIn()
  }
}

func (bstnode *BstNode) PrintPos(){
  if bstnode.left != nil{
    bstnode.left.PrintIn()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintIn()
  }
  fmt.Print(", ", bstnode.value)
}

func (bstnode *BstNode) Max() int{
  if bstnode.rigth != nil{
    return bstnode.rigth.Max()
  }else{
    return bstnode.value
  }
}

func (bstnode *BstNode) Min() int{
  if bstnode.left != nil{
    return bstnode.left.Min()
  }else{
    return bstnode.value
  }
}

func (bstnode *BstNode) Remove(value int){
  
}

func (bstnode *BstNode) Altura() int{
  return 1
}

func CreateBst(v []int) *BstNode{
  repete := IsRepet(v)
  if repete == false{
    root := NewNode(v[0])
    for i:=1; i<len(v); i++{
      root.Add(v[i])
    }
    return root
  }else{
    return nil
  }
}

func IsRepet(v []int) bool{
  repete := false
  for i:=0; i<len(v); i++{
    for j:=i+1; j<len(v)-1; j++{
      if v[i] == v[j]{
        repete = true
        break
      }
    }
  }
  return repete
}

func (bstNode *BstNode) isBst() bool{
  return false
}