package main

import (
	"fmt"
	//"main/packs/sort"
	"main/packs/tree"
	//"math/rand"
	//"time"
)

func main() {
	arvore := tree.NewNode(0)

	for i := 1; i < 250; i++ {
		arvore = arvore.Add(i)
	}

	fmt.Println("altura da arvore = ", arvore.Get_height())
	arvore.PrintPre()

	// 	arvore := tree.NewNode(20)
	// 	arvore = arvore.Add(29)
	// 	arvore = arvore.Add(35)
	// 	arvore = arvore.Add(51)
	// 	arvore = arvore.Add(13)
	// 	arvore = arvore.Add(25)
	// 	arvore = arvore.Add(18)
	// 	arvore = arvore.Add(19)
	// 	arvore.PrintPre()
	//arv_rotdir.PrintPre()

	/*
		v := []int{10, 7, 3, 8, 15, 20}
		t := tree.CreateBst(v)
		u := tree.BstNode{}
		u.Add(10)
		u.Add(7)
		u.Add(3)
		u.Add(8)
		u.Add(20)
		t.Remove(3)
		t.Remove(20)
		u.Remove(3)
		fmt.Println("size da arvore = ", t.Size())
		t.PrintIn()
		fmt.Println("size da arvore = ", u.Size())
		u.PrintIn()*/

	/* testes algortitmos de ordenação
	vetor := FullArrayDecres(10000)
	fmt.Println("vetor = ", vetor)
	fmt.Println()

	//BubbleSort
	start := time.Now()
	vetor_bubble := sort.BubbleSort(vetor)
	end := time.Since(start)
	fmt.Println("vetor ordenado com BubbleSort = ", vetor_bubble)
	fmt.Println()

	//SelectionSort
	start1 := time.Now()
	vetor_selection := sort.SelectionSort(vetor)
	end1 := time.Since(start1)
	fmt.Println("vetor ordenado com SelectionSort = ", vetor_selection)

	//InsertionSort
	start2 := time.Now()
	vetor_insertion := sort.InsertionSort(vetor)
	end2 := time.Since(start2)
	fmt.Println("vetor ordenado com InsertionSort = ", vetor_insertion)

	//MergeSort
	start3 := time.Now()
	vetor_merge := sort.MergeSort(vetor,len(vetor))
	end3 := time.Since(start3)
	fmt.Println("vetor ordenado com MergeSort = ", vetor_merge)

	vetor_quick := vetor

	//QuickSort
	start4 := time.Now()
	sort.QuickSort(vetor_quick,0,len(vetor_quick)-1)
	end4 := time.Since(start4)
	fmt.Println("vetor ordenado com QuickSort = ", vetor_quick)

	//CountingSort
	start5 := time.Now()
	vetor_couting := sort.CountingSort(vetor)
	end5 := time.Since(start5)
	fmt.Println("vetor ordenado com CountingSort = ", vetor_couting)
	fmt.Println()

	fmt.Println("tempo gasto com BubbleSort O(n²) = ", end)
	fmt.Println("tempo gasto com SelectionSort O(n²) = ", end1)
	fmt.Println("tempo gasto com InsertionSort O(n²) = ", end2)
	fmt.Println("tempo gasto com MergeSort O(logn) = ", end3)
	fmt.Println("tempo gasto com QuickSort O(nlogn) = ", end4)
	fmt.Println("tempo gasto com CountingSort O(n) = ", end5)


	// fazendo testes em arvores
	// vetor := []string{"c","a","r","l","o","s"}
	// t := tree.CreateBst_str(vetor)
	// // bst_verify := t.IsBst()

	// if t != nil{
	//   t.PrintPos()
	//   // fmt.Println(bst_verify)
	// }else{
	//   fmt.Println(t)
	// }*/
}

/*
func Add_numeros(qtd int, intervalo int) {
	tree := tree.BstNode{}
	vetor_numeros := FullArrayRandom(qtd, intervalo)

	for i := 0; i < len(vetor_numeros); i++ {
		tree.Add(vetor_numeros[i])
	}
	TestTree(&tree, vetor_numeros)
	fmt.Print("print In = ")
	tree.PrintIn()
	fmt.Println()
	fmt.Print("print Pre = ")
	tree.PrintPre()
	fmt.Println()
	fmt.Print("print Pos = ")
	tree.PrintPos()
	fmt.Println()
	max := tree.Max()
	min := tree.Min()
	fmt.Println("max = ", max, " min = ", min)
	fmt.Println()

}

func TestTree(tree *tree.BstNode, vetor []int) {
	fmt.Println("testando a busca na árvore = ", vetor)

	for i := 0; i < 10; i++ {
		busca := tree.Search(i)
		fmt.Println("procurando o ", i, "achou? ", busca)
	}
	fmt.Println()
}

func FullArrayRandom(tam int, intervalo int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = rand.Intn(intervalo)
	}
	return vetor
}

// func FullArraySemRepet(tam int) []int{

// }

func FullArrayCres(tam int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = i
	}
	return vetor
}

func FullArrayDecres(tam int) []int {
	vetor := make([]int, tam)

	for i := 0; i < tam; i++ {
		vetor[i] = (tam - 1) - i
	}
	return vetor
}

package main

import (
	"fmt"
)

type ITree interface {
	Add(value int) *BstNode
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	Height() int
	Remove(value int) *BstNode
	IsBst() bool
	Size() int
	RotRight() *BstNode
	RotLeft() *BstNode
	Rebalance() *BstNode
	UpdateProperties()
}

type BstNode struct {
	left   *BstNode
	value  int
	height int
	bf     int
	right  *BstNode
}

func NewNode(value int) *BstNode {
	node := BstNode{}
	node.value = value
	node.right = nil
	node.left = nil
	node.height = 0
	node.bf = 0
	return &node
}

func (bstNode *BstNode) Add(value int) {
	if value <= bstNode.value { //insercao à esq
		if bstNode.left == nil {
			bstNode.left = NewNode(value)
		} else {
			bstNode.left.Add(value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = NewNode(value)
		} else {
			bstNode.right.Add(value)
		}
	}
	//bstNode.UpdateProperties()
	//return bstNode.Rebalance()
}

/*func (bstNode *BstNode) Search(value int) bool {
	if value == bstNode.value {
		return true
	} else if value < bstNode.value && bstNode.left != nil {
		return bstNode.left.Search(value)
	} else if value > bstNode.value && bstNode.right != nil {
		return bstNode.right.Search(value)
	} else {
		return false
	}
}

func (bstNode *BstNode) Min() int {
	if bstNode.left == nil {
		return bstNode.value
	} else {
		return bstNode.left.Min()
	}
}

func (bstNode *BstNode) Max() int {
	if bstNode.right == nil {
		return bstNode.value
	} else {
		return bstNode.right.Max()
	}
}

func (bstNode *BstNode) PrintPre() {
	fmt.Print(bstNode.value, ", ")
	if bstNode.left != nil {
		bstNode.left.PrintPre()
	}
	if bstNode.right != nil {
		bstNode.right.PrintPre()
	}
}

func (bstNode *BstNode) PrintIn() {
	if bstNode.left != nil {
		bstNode.left.PrintIn()
	}
	fmt.Print(bstNode.value, ", ")
	if bstNode.right != nil {
		bstNode.right.PrintIn()
	}
}

func (bstNode *BstNode) PrintPos() {
	if bstNode.left != nil {
		bstNode.left.PrintPos()
	}
	if bstNode.right != nil {
		bstNode.right.PrintPos()
	}
	fmt.Print(bstNode.value, ", ")
}

func (bstNode *BstNode) Height() int {
	hBasedOnLeft := 0
	hBasedOnRight := 0
	if bstNode.left != nil {
		hBasedOnLeft = 1 + bstNode.left.Height()
	}
	if bstNode.right != nil {
		hBasedOnRight = 1 + bstNode.right.Height()
	}
	if hBasedOnLeft >= hBasedOnRight {
		return hBasedOnLeft
	} else {
		return hBasedOnRight
	}
}
/*
func (bstNode *BstNode) Remove(value int) *BstNode {
	if value < bstNode.value {
		bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value {
		bstNode.right = bstNode.right.Remove(value)
	} else { //encontramos o nó a ser removido
		if bstNode.left == nil && bstNode.right == nil { //caso 1: nó folha
			return nil
		} else if bstNode.left != nil && bstNode.right == nil { //caso 2: nó com 1 filho (à esquerda)
			return bstNode.left
		} else if bstNode.left == nil && bstNode.right != nil { //caso 2: nó com 1 filho (à direita)
			return bstNode.right
		} else { //caso 3: nó com 2 filhos
			maxEsq := bstNode.left.Max()
			bstNode.value = maxEsq
			bstNode.left = bstNode.left.Remove(maxEsq)
			return bstNode
		}
	}
	bstNode.UpdateProperties()
	return bstNode.Rebalance()
}

func (bstNode *BstNode) IsBst() bool {
	if bstNode.left != nil {
		if bstNode.value > bstNode.left.value {
			return bstNode.left.IsBst()
		} else {
			return false
		}
	}
	if bstNode.right != nil {
		if bstNode.value < bstNode.right.value {
			return bstNode.right.IsBst()
		} else {
			return false
		}
	}
	return true
}

func (bstNode *BstNode) RotRight() *BstNode {
  left := bstNode.left
	bstNode.left = left.right
	left.right = bstNode
	//bstNode.UpdateProperties()
	//left.UpdateProperties()
	return left
}

func (bstNode *BstNode) RotLeft() *BstNode {
	right := bstNode.right
	bstNode.right = right.left
	right.left = bstNode
	//bstNode.UpdateProperties()
	//right.UpdateProperties()
	return right
}

// Update Height and BalanceFactor
func (bstNode *BstNode) UpdateProperties() {
	//atualizar altura
	heightRight := 0
	heightLeft := 0
	if bstNode.right == nil && bstNode.left == nil {
		bstNode.height = 0
	} else {
		if bstNode.right != nil {
			heightRight = bstNode.right.height
		}
		if bstNode.left != nil {
			heightLeft = bstNode.left.height
		}
		if heightRight > heightLeft {
			bstNode.height = 1 + heightRight
		} else {
			bstNode.height = 1 + heightLeft
		}
	}
	bstNode.bf = heightRight - heightLeft
}

func (bstNode *BstNode) Rebalance() *BstNode {
	if bstNode.bf <= -2 { //LEFT
		if bstNode.left.bf == -1 { //LEFT-LEFT
			return bstNode.RebalanceLeftLeft()
		} else if bstNode.left.bf == 0 { //LEFT-NEUTRAL
			return bstNode.RebalanceLeftLeft()
		} else { //LEFT-RIGHT
			return bstNode.RebalanceLeftRight()
		}
	} else if bstNode.bf >= 2 { //RIGHT
		if bstNode.right.bf == 1 { //RIGHT-RIGHT
			return bstNode.RebalanceRightRight()
		} else if bstNode.right.bf == 0 { //RIGHT-NEUTRAL
			return bstNode.RebalanceRightRight()
		} else { //RIGHT-LEFT
			return bstNode.RebalanceRightLeft()
		}
	}
	return bstNode
}

func (bstNode *BstNode) RebalanceRightRight() *BstNode {
	return bstNode.RotLeft()
}

func (bstNode *BstNode) RebalanceRightLeft() *BstNode {
	bstNode.right = bstNode.right.RotRight()
	return bstNode.RebalanceRightRight()
}

func (bstNode *BstNode) CheckBF() bool {
	if bstNode.bf > -2 && bstNode.bf < 2 {
		if bstNode.left != nil {
			return bstNode.left.CheckBF()
		}
		if bstNode.right != nil {
			return bstNode.right.CheckBF()
		}
		return true
	} else {
		return false
	}
}

func main() {
  arvore := BstNode{}
  arvore.Add(5)
  arvore.Add(2)
  arvore.Add(15)
  arvore.Add(11)
  arvore.Add(25)
  arvore.PrintPre()
  arvore_rot := arvore.RotRight()
  arvore_rot.PrintPre()
}*/
