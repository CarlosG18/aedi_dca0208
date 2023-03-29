package list

import ("fmt"
  "errors"
)

type ArrayList struct {
	values []int
	tam    int
}

func (arraylist *ArrayList) Init(capacity int) error{
  if capacity > 0{
    arraylist.values = make([]int, capacity)
    return nil
  }else{
    return errors.New("não é possivel criar um vetor com capacidade negativa")
  }
}

func (arraylist *ArrayList) RemoveOnIndex(index int) error{
if arraylist.tam > 0{
  if index < arraylist.tam && index >= 0{
		for i := index; i < arraylist.tam-1; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.tam--
    arraylist.Print("RemoveOnIndex")
    return nil
  }else{
    return errors.New("index não aceito!")
  }
}else{
  return errors.New("a lista não possui elementos")
}
}

func (arraylist *ArrayList) Remove() error{
	if arraylist.tam > 0 {
		arraylist.tam--
    arraylist.Print("Remove")
    return nil
	}else{
    return errors.New("não é possivel remover um vetor sem elementos")
  }
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) error{
  if index > arraylist.tam || index < 0{
    if arraylist.tam == len(arraylist.values) {
		arraylist.Double()
	  }
    for i := arraylist.tam; i < index; i-- {
		arraylist.values[i] = arraylist.values[i-1]
	  }
	  arraylist.values[index] = value
	  arraylist.tam++
    arraylist.Print("AddOnIndex")
    return nil
  }else{
    return errors.New("index não aceito!")
  }
}

func (arraylist *ArrayList) Double() {
	doublelist := make([]int, len(arraylist.values)*2)
	for i := 0; i < arraylist.tam; i++ {
		doublelist[i] = arraylist.values[i]
	}
	arraylist.values = doublelist
}

func (arraylist *ArrayList) Get(index int) (int,error) {
	if index >= 0 && index < arraylist.tam {
		return arraylist.values[index], nil
  }
  return -1, errors.New("index não aceito!")
}

func (arraylist *ArrayList) Set(index int, value int) error{
	if index >= 0 && index < arraylist.tam {
		arraylist.values[index] = value
    arraylist.Print("Set")
    return nil
	}
  arraylist.Print("Set")
  return errors.New("index não aceito!")
}

func (arraylist *ArrayList) Add(value int) {
	if arraylist.tam == len(arraylist.values) {
		arraylist.Double()
	}
	arraylist.values[arraylist.tam] = value
	arraylist.tam++
  arraylist.Print("Add")
}

func (arraylist *ArrayList) Size() int {
	return arraylist.tam
}

func (arraylist *ArrayList) Print(operation string){
  fmt.Println("operação = ", operation)
if arraylist.tam > 0 {
  fmt.Print("[")
  for i := 0; i < arraylist.tam; i++ {
    if i > 0 {
      fmt.Print(",")
    }
    fmt.Print(arraylist.values[i])
  }
  fmt.Print("]")
}else{
  fmt.Println("o vetor não possui valores")
}
  fmt.Println()
  fmt.Println("tamanho do vetor = ", arraylist.Size())
  fmt.Println()
}