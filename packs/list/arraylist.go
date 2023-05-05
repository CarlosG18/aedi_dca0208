package list

import (
  "errors"
)

type ArrayList struct {
	values []int
	tam    int
}

func (arraylist *ArrayList) Init(capacity int) error{
  if capacity > 0{
    arraylist.values = make([]int, capacity)
    arraylist.tam = 0
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
    return nil
	}else{
    return errors.New("não é possivel remover um vetor sem elementos")
  }
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) error{
  if index > arraylist.tam || index < 0{
    return errors.New("index não aceito!")
  }else{
    if arraylist.tam == len(arraylist.values) {
		arraylist.Double()
	  }
    for i := arraylist.tam; i > index; i-- {
		arraylist.values[i] = arraylist.values[i-1]
	  }
	  arraylist.values[index] = value
	  arraylist.tam++
    return nil
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

func (arraylist *ArrayList) Set(value int, index int) error{
	if index >= 0 && index <= arraylist.tam {
		arraylist.values[index] = value
    return nil
	}
  return errors.New("index não aceito!")
}

func (arraylist *ArrayList) Add(value int) {
	if arraylist.tam == len(arraylist.values) {
		arraylist.Double()
	}
	arraylist.values[arraylist.tam] = value
	arraylist.tam++
}

func (arraylist *ArrayList) Size() int {
	return arraylist.tam
}