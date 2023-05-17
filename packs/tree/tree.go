package tree

type ITREE interface{
  Add(value int)
  Busca(value int) bool
}