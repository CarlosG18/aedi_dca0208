import tree

type Bst struct{
  root *BstNode
  size int
}

type BstNode struct{
  left *BstNode
  value int
  rigth *BstNode
}

func (bstnode *BstNode) NewNode(value int) return *BstNode{
  node = {}BstNode
  bstnode.left = nil
  bstnode.rigth = nil
  bstnode.value = value
  return &node
}

func (bstnode *BstNode) Add(value int){
  
}

