package tree

type ITREE interface{
  Add(value int)
  Busca(value int) bool
  Min() int
  Max() int
  PrintPre()
  PrintIn()
  PrintPos()
  CreateBst(v []int) *BstNode
}