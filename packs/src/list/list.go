package list

type IList interface{
  Add(value int)
  Remove() 
  RemoveOnIndex(index int) error
  AddOnIndex(value int, index int) error
  Size() (int, error)
  Get(index int) (int, error)
  Set(value int, index int) error
  Print(operation string)
}