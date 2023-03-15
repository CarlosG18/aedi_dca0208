package list

import "fmt"

type ArrayList struct {
	values []int
	tam    int
}

func (arraylist *ArrayList) Init(capacity int) {
	arraylist.values = make([]int, capacity)
}

func (arraylist *ArrayList) RemoveOnIndex(index int) {
	if index < arraylist.tam || index >= 0 {
		for i := index; i < arraylist.tam-1; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.tam--
	}
}

func (arraylist *ArrayList) Remove() {
	if arraylist.tam > 0 {
		arraylist.tam--
	}
	//tratamento do erro

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

func (arraylist *ArrayList) Print(){
  fmt.Println("[ ")
  for i := 0; i < len(arraylist.values); i++{
    fmt.Println(arraylist.values[i])
fmt.Println(" ")

  }
  fmt.Println("]")
}