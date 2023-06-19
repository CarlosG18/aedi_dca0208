package tree

import "fmt"

type BstNode struct {
	left   *BstNode
	value  int
	bf     int
	height int
	rigth  *BstNode
}

func (bstnode *BstNode) Teste() {
	fmt.Println("antes: ")
	bstnode.PrintPre()
	bstnode = bstnode.LeftRight()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("depois: ")
	bstnode.PrintPre()
}

func NewNode(value int) *BstNode {
	node := BstNode{}
	node.value = value
	node.bf = 0
	node.height = 0
	return &node
}

func (bstnode *BstNode) Get_height() int {
	return bstnode.height
}

func (bstnode *BstNode) UpdateProp() {
	alt_dir := 0
	alt_esq := 0
	if bstnode.left != nil {
		alt_esq = bstnode.left.Altura()
	}
	if bstnode.rigth != nil {
		alt_dir = bstnode.rigth.Altura()
	}
	bstnode.bf = alt_dir - alt_esq
	bstnode.height = bstnode.Altura()
}

func (bstnode *BstNode) RotLeft() *BstNode {
	rigth := bstnode.rigth
	bstnode.rigth = rigth.left
	rigth.left = bstnode
	bstnode.UpdateProp()
	rigth.UpdateProp()
	return rigth
}

func (bstnode *BstNode) LeftLeft() *BstNode {
	return bstnode.RotRight()
}

func (bstnode *BstNode) Left_Neutral() *BstNode {
	return bstnode.RotRight()
}

func (bstnode *BstNode) LeftRight() *BstNode {
	left := bstnode.left
	bstnode.left = left.RotLeft()
	return bstnode.RotRight()
}

func (bstnode *BstNode) RotRight() *BstNode {
	left := bstnode.left
	bstnode.left = left.rigth
	left.rigth = bstnode
	bstnode.UpdateProp()
	left.UpdateProp()
	return left
}

func (bstnode *BstNode) Right_Neutral() *BstNode {
	return bstnode.RotLeft()
}

func (bstnode *BstNode) RightRight() *BstNode {
	return bstnode.RotLeft()
}

func (bstnode *BstNode) RightLeft() *BstNode {
	rigth := bstnode.rigth
	bstnode.rigth = rigth.RotRight()
	return bstnode.RotLeft()
}

/*  fazer ajustes */
func (bstnode *BstNode) Add(value int) *BstNode {
	if value > bstnode.value {
		if bstnode.rigth != nil {
			bstnode.rigth = bstnode.rigth.Add(value)
		} else {
			bstnode.rigth = NewNode(value)
		}
	} else {
		if bstnode.left != nil {
			bstnode.left = bstnode.left.Add(value)
		} else {
			bstnode.left = NewNode(value)
		}
	}
	bstnode.UpdateProp()
	return bstnode.Rebalance()
}

func (bstnode *BstNode) Rebalance() *BstNode {
	if bstnode.bf <= -2 { //LEFT
		if bstnode.left.bf == 0 {
			return bstnode.Left_Neutral()
		} else if bstnode.left.bf == -1 {
			return bstnode.LeftLeft()
		} else {
			return bstnode.LeftRight()
		}
	} else if bstnode.bf >= 2 { //RIGHT
		if bstnode.rigth.bf == 0 {
			return bstnode.Right_Neutral()
		} else if bstnode.rigth.bf == -1 {
			return bstnode.RightLeft()
		} else {
			return bstnode.RightRight()
		}
	} else {
		return bstnode
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

// navegação em arvores
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
	fmt.Print(", ", bstnode.value, " bf = ", bstnode.bf, " altura = ", bstnode.height)
	fmt.Println()

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
	alt_base_dir := 0
	alt_base_esq := 0
	if bstnode.left != nil {
		alt_base_esq = 1 + bstnode.left.Altura()
	}
	if bstnode.rigth != nil {
		alt_base_dir = 1 + bstnode.rigth.Altura()
	}
	if alt_base_dir > alt_base_esq {
		return alt_base_dir
	} else {
		return alt_base_esq
	}
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
