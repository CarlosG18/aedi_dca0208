package hash

import 
  "fmt"

type Tupla struct {
	key   int
	value string
}

type TableHash struct {
	bucker   [][]Tupla
	factor_carga float32
	tamTable int
}

func Init(n int) *TableHash{
  table := TableHash{}
  table.bucker = make([][]Tupla,n)
  table.tamTable = 0
  return &table
}

var empty_tupla = Tupla{}

func (table *TableHash) Add(key int, value string) {
  if table.factor_carga < 0.75{
    index_bucker := key % len(table.bucker)
    tam_index_bucker := len(table.bucker[index_bucker])
    table.bucker[index_bucker] = append(table.bucker[index_bucker], Tupla{})
    table.bucker[index_bucker][tam_index_bucker].key = key
    table.bucker[index_bucker][tam_index_bucker].value = value
    table.tamTable++
    table.Factor_carga()
  }else{
    new_table := Init(8*len(table.bucker))
    
  }
}

func (table *TableHash) Remove(key int) {
  is_key := table.Search(key)
  if is_key{
    index_bucker := key % len(table.bucker)
    for i:=0; i<len(table.bucker[index_bucker]); i++{
      if table.bucker[index_bucker][i].key == key{
        table.bucker[index_bucker] = append(table.bucker[index_bucker][:i], table.bucker[index_bucker][i+1:]...)
      }
    }
  }else{
    fmt.Println("não é possivel remover pois essa chave não existe!")
  }
  table.tamTable--
  table.factor_carga()
}

func (table *TableHash) Get(key int) *Tupla{
  tupla := Tupla{}
  is_key := table.Search(key)
  
  if is_key{
    index_bucker := key % len(table.bucker)
    for i:=0; i<len(table.bucker[index_bucker]); i++{
      if table.bucker[index_bucker][i].key == key{
        tupla.key = table.bucker[index_bucker][i].key
        tupla.value = table.bucker[index_bucker][i].value
        break
      }
    }
  }
  return &tupla
}

func (table *TableHash) Search(key int) bool {
  index_bucker := key % len(table.bucker)
  for i:=0; i<len(table.bucker[index_bucker]); i++{
    if table.bucker[index_bucker][i].key == key{
      return true
      break
    }
  }
  return false
}

func (table *TableHash) Hash_function(key int) int {
	return 1
}

func (table *TableHash) Print_table(){
	for i:=0; i < len(table.bucker); i++{
	  for j:=0; j < len(table.bucker[i]); j++{
	    fmt.Println("indicie[", i, "] = { ", table.bucker[i][j].key, ", ",table.bucker[i][j].value, " }")
	  }
	}
	factor := table.Factor_carga()
	fmt.Println("fator de carga = ", factor)
	fmt.Println()
}

func (table *TableHash) Factor_carga() {
	table.factor_carga = float32(table.tamTable)/float32(len(table.bucker))
}