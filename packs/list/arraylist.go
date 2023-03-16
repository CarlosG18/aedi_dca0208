package list

import "fmt"

type ArrayList struct {
	values []int
	tam    int
}

func (arraylist *ArrayList) Init(capacity int) {
	arraylist.values = make([]int, capacity)
  arraylist.tam = 0
}

func (arraylist *ArrayList) RemoveOnIndex(index int) {
	if index < arraylist.tam && index >= 0 && arraylist.tam > 0{
		for i := index; i < arraylist.tam-1; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.tam--
  }else{
    fmt.Print("index not encontry")
  }
  arraylist.Print("RemoveOnIndex")
}

func (arraylist *ArrayList) Remove() {
	if arraylist.tam > 0 {
		arraylist.tam--
	}
	//tratamento do erro
  arraylist.Print("Remove")
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) {
	if arraylist.tam == len(arraylist.values) {
		arraylist.Double()
	}

	for i := arraylist.tam; i < index; i-- {
		arraylist.values[i] = arraylist.values[i-1]
	}
	arraylist.values[index] = value
	arraylist.tam++
  arraylist.Print("AddOnIndex")
}

func (arraylist *ArrayList) Double() {
	doublelist := make([]int, len(arraylist.values)*2)
	//copiar os elementos
	for i := 0; i < arraylist.tam; i++ {
		doublelist[i] = arraylist.values[i]
	}
	arraylist.values = doublelist
}

func (arraylist *ArrayList) Get(index int) int {
	if index >= 0 && index < arraylist.tam {
		return arraylist.values[index]
	}

	return -1
}

func (arraylist *ArrayList) Set(index int, value int) {
	if index >= 0 && index < arraylist.tam {
		arraylist.values[index] = value
	}
  arraylist.Print("Set")
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