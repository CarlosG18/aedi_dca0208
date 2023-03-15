package list

type Ilist interface{
  Init()
  Add(value int)
  Remove()
  RemoveOnIndex(index int)
  AddOnIndex(value int, index int)
  Size()
  Double()
  Get(index int)
  Set(value int, index int)
}