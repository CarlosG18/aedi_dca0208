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
    bstnode.left.PrintPre()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintPre()
  }
}

func (bstnode *BstNode) PrintPos(){
  if bstnode.left != nil{
    bstnode.left.PrintPos()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintPos()
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

func (bstNode *BstNode) IsBst() bool{
   value := bstNode.value
  if bstNode.left != nil{
    if bstNode.left.value > value{
      return false
    }else{
      bstNode.left.IsBst()
    }  
  }
  if bstNode.rigth != nil{
    if bstNode.rigth.value < value{
      return false
    }else{
      bstNode.rigth.IsBst()
    }  
  }
  return true
}

// versao string

type BstNode_str struct{
  left *BstNode_str
  value string
  rigth *BstNode_str
}

func NewNode_str(value string) *BstNode_str{
  node := BstNode_str{}
  node.left = nil
  node.rigth = nil
  node.value = value
  return &node
}

func (bstnode *BstNode_str) Add(value string){
  if value > bstnode.value{
    if bstnode.rigth != nil{
      bstnode.rigth.Add(value)
    }else{
      bstnode.rigth = NewNode_str(value)
    }
  }else{
    if bstnode.left != nil{
      bstnode.left.Add(value)
    }else{
      bstnode.left = NewNode_str(value)
    }
  }
}

func (bstnode *BstNode_str) PrintIn(){
  if bstnode.left != nil{
    bstnode.left.PrintIn()
  }
  fmt.Print(bstnode.value, ", ")
  if bstnode.rigth != nil{
    bstnode.rigth.PrintIn()
  }
}

func (bstnode *BstNode_str) PrintPre(){
  fmt.Print(", ", bstnode.value)
  if bstnode.left != nil{
    bstnode.left.PrintPre()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintPre()
  }
}

func (bstnode *BstNode_str) PrintPos(){
  if bstnode.left != nil{
    bstnode.left.PrintPos()
  }
  if bstnode.rigth != nil{
    bstnode.rigth.PrintPos()
  }
  fmt.Print(", ", bstnode.value)
}

func CreateBst_str(v []string) *BstNode_str{
  repete := IsRepet_str(v)
  if repete == false{
    root := NewNode_str(v[0])
    for i:=1; i<len(v); i++{
      root.Add(v[i])
    }
    return root
  }else{
    return nil
  }
}

func IsRepet_str(v []string) bool{
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

func (bstnode *BstNode_str) Search(value string) bool{
  
  
  if value < bstnode.value && bst{
    bstnode.left.Search(value)
  }else if value > bstnode.value{
    bstnode.rigth.Search(value)
  }else{
    return true
  }

  return false
}