package tree

import "fmt"

type BstNode struct {
	left  *BstNode
	value int
	bf int
	height int
	rigth *BstNode
}

func NewNode(value int) *BstNode {
	node := BstNode{}
	node.value = value
	node.bf = 0
	node.height = 0
	return &node
}

func (bstnode *BstNode) UpdateProp(){
  //update propriedades 
}

func (bstnode *BstNode) RotLeft() *BstNode{
  rigth := bstnode.rigth
  bstnode.rigth = rigth.left
  rigth.left = bstnode
  return rigth
}

func (bstnode *BstNode) RotRight() *BstNode{
  left := bstnode.left
  if left.rigth != nil{
    bstnode.left = left.rigth
  }else{
    bstnode.left = nil
  }
  left.rigth = bstnode
  return left
}

/*  fazer ajustes */
func (bstnode *BstNode) Add(value int) {
	if value > bstnode.value {
		if bstnode.rigth != nil {
			bstnode.rigth.Add(value)
		} else {
			bstnode.rigth = NewNode(value)
		}
	} else {
		if bstnode.left != nil {
			bstnode.left.Add(value)
		} else {
			bstnode.left = NewNode(value)
		}
	}
}

func (bstnode *BstNode) Search(value int) bool {
	if value == bstnode.value {
		return true
	} else if value > bstnode.value {
		if bstnode.rigth == nil {
			return false
		} else {
			return bstnode.rigth.Search(value)
		}
	} else {
		if bstnode.left == nil {
			return false
		} else {
			return bstnode.left.Search(value)
		}
	}
}

//navegação em arvores
func (bstnode *BstNode) PrintIn() {
	if bstnode.left != nil {
		bstnode.left.PrintIn()
	}
	fmt.Print(bstnode.value, ", ")
	if bstnode.rigth != nil {
		bstnode.rigth.PrintIn()
	}
}

func (bstnode *BstNode) PrintPre() {
	fmt.Print(", ", bstnode.value)
	if bstnode.left != nil {
		bstnode.left.PrintPre()
	}
	if bstnode.rigth != nil {
		bstnode.rigth.PrintPre()
	}
}

func (bstnode *BstNode) PrintPos() {
	if bstnode.left != nil {
		bstnode.left.PrintPos()
	}
	if bstnode.rigth != nil {
		bstnode.rigth.PrintPos()
	}
	fmt.Print(", ", bstnode.value)
}

func (bstnode *BstNode) Max() int {
	if bstnode.rigth != nil {
		return bstnode.rigth.Max()
	} else {
		return bstnode.value
	}
}

func (bstnode *BstNode) Min() int {
	if bstnode.left != nil {
		return bstnode.left.Min()
	} else {
		return bstnode.value
	}
}

func (bstNode *BstNode) Remove(value int) *BstNode {
	if value < bstNode.value {
		bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value {
		bstNode.rigth = bstNode.rigth.Remove(value)
	} else { //encontramos o nó a ser removido
		if bstNode.left == nil && bstNode.rigth == nil { //caso 1: nó folha
			return nil
		} else if bstNode.left != nil && bstNode.rigth == nil { //caso 2: nó com 1 filho (à esquerda)
			return bstNode.left
		} else if bstNode.left == nil && bstNode.rigth != nil { //caso 2: nó com 1 filho (à direita)
			return bstNode.rigth
		} else { //caso 3: nó com 2 filhos
			maxEsq := bstNode.left.Max()
			bstNode.value = maxEsq
			bstNode.left = bstNode.left.Remove(maxEsq)
			return bstNode
		}
	}
	return bstNode
}

func (bstnode *BstNode) Altura() int {
	return 1
}

func (bstnode *BstNode) Size() int {
	size := 1

	if bstnode.left != nil {
		size += bstnode.left.Size()
	}
	if bstnode.rigth != nil {
		size += bstnode.rigth.Size()
	}
	return size
}

func CreateBst(v []int) *BstNode {
	repete := IsRepet(v)
	if repete == false {
		root := NewNode(v[0])
		for i := 1; i < len(v); i++ {
			root.Add(v[i])
		}
		return root
	} else {
		return nil
	}
}

func IsRepet(v []int) bool {
	repete := false
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v)-1; j++ {
			if v[i] == v[j] {
				repete = true
				break
			}
		}
	}
	return repete
}

func (bstNode *BstNode) IsBst() bool {
	value := bstNode.value
	if bstNode.left != nil {
		if bstNode.left.value > value {
			return false
		} else {
			bstNode.left.IsBst()
		}
	}
	if bstNode.rigth != nil {
		if bstNode.rigth.value < value {
			return false
		} else {
			bstNode.rigth.IsBst()
		}
	}
	return true
}